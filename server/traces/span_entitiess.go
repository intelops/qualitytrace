package traces

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/intelops/qualitytrace/agent/workers/trigger"
	"github.com/intelops/qualitytrace/server/pkg/timing"
	"go.opentelemetry.io/otel/trace"
)

const (
	QualitytraceMetadataFieldStartTime         string = "qualitytrace.span.start_time"
	QualitytraceMetadataFieldEndTime           string = "qualitytrace.span.end_time"
	QualitytraceMetadataFieldDuration          string = "qualitytrace.span.duration"
	QualitytraceMetadataFieldType              string = "qualitytrace.span.type"
	QualitytraceMetadataFieldName              string = "qualitytrace.span.name"
	QualitytraceMetadataFieldParentID          string = "qualitytrace.span.parent_id"
	QualitytraceMetadataFieldKind              string = "qualitytrace.span.kind"
	QualitytraceMetadataFieldStatusCode        string = "qualitytrace.span.status_code"
	QualitytraceMetadataFieldStatusDescription string = "qualitytrace.span.status_description"
)

func NewAttributes(inputs ...map[string]string) Attributes {
	attr := Attributes{
		mutex:  &sync.Mutex{},
		values: make(map[string]string),
	}

	for _, input := range inputs {
		for key, value := range input {
			attr.values[key] = value
		}
	}

	return attr
}

type Attributes struct {
	mutex  *sync.Mutex
	values map[string]string
}

func (a Attributes) Values() map[string]string {
	a.lock()
	defer a.unlock()

	m := make(map[string]string, len(a.values))
	for key, value := range a.values {
		m[key] = value
	}
	return m
}

func (a Attributes) Len() int {
	return len(a.values)
}

func (a *Attributes) lock() {
	if a.mutex == nil {
		a.mutex = &sync.Mutex{}
	}

	a.mutex.Lock()
}

func (a *Attributes) unlock() {
	if a.mutex != nil {
		a.mutex.Unlock()
	}
}

func (a Attributes) MarshalJSON() ([]byte, error) {
	a.lock()
	defer a.unlock()

	return json.Marshal(a.values)
}

func (a *Attributes) UnmarshalJSON(in []byte) error {
	if a.mutex == nil {
		a.mutex = &sync.Mutex{}
	}
	a.lock()
	defer a.unlock()

	a.values = make(map[string]string, 0)

	return json.Unmarshal(in, &a.values)
}

func (a Attributes) GetExists(key string) (string, bool) {
	a.lock()
	defer a.unlock()

	if v, ok := a.values[key]; ok {
		return v, true
	}

	return "", false
}

func (a Attributes) Get(key string) string {
	v, _ := a.GetExists(key)
	return v
}

func (a Attributes) Set(key, value string) Attributes {
	a.lock()
	defer a.unlock()
	a.values[key] = value

	return a
}

func (a Attributes) Delete(key string) {
	a.lock()
	defer a.unlock()

	delete(a.values, key)
}

func (a Attributes) SetPointerValue(key string, value *string) {
	if value != nil {
		a.values[key] = *value
	}
}

type Spans []Span

func (s Spans) ForEach(fn func(ix int, _ Span) bool) Spans {
	for i, span := range s {
		doNext := fn(i, span)
		if !doNext {
			break
		}
	}
	return s
}

func (s Spans) OrEmpty(fn func()) Spans {
	if len(s) == 0 {
		fn()
	}
	return s
}

type SpanKind string

var (
	SpanKindClient       SpanKind = "client"
	SpanKindServer       SpanKind = "server"
	SpanKindConsumer     SpanKind = "consumer"
	SpanKindProducer     SpanKind = "producer"
	SpanKindInternal     SpanKind = "internal"
	SpanKindUnespecified SpanKind = "unespecified"
)

type Span struct {
	ID         trace.SpanID
	Name       string
	StartTime  time.Time
	EndTime    time.Time
	Attributes Attributes
	Kind       SpanKind
	Events     []SpanEvent
	Status     *SpanStatus

	Parent   *Span   `json:"-"`
	Children []*Span `json:"-"`
}

type SpanStatus struct {
	Code        string
	Description string
}

func (s *Span) injectEventsIntoAttributes() {
	if s.Events == nil {
		s.Events = make([]SpanEvent, 0)
	}

	eventsJson, _ := json.Marshal(s.Events)
	s.Attributes.Set("span.events", string(eventsJson))
}

type SpanEvent struct {
	Name       string     `json:"name"`
	Timestamp  time.Time  `json:"timestamp"`
	Attributes Attributes `json:"attributes"`
}

type encodedSpan struct {
	ID         string
	Name       string
	Kind       string
	StartTime  string
	EndTime    string
	Attributes Attributes
	Children   []encodedSpan
}

const nilSpanID = "0000000000000000"

func (es encodedSpan) isValidID() bool {
	if es.ID == nilSpanID || es.ID == "" {
		return false
	}
	return true
}

func (s Span) IsZero() bool {
	return !s.ID.IsValid()
}

func (s Span) MarshalJSON() ([]byte, error) {
	enc := encodeSpan(s)
	return json.Marshal(&enc)
}

func encodeSpan(s Span) encodedSpan {
	return encodedSpan{
		ID:         s.ID.String(),
		Name:       s.Name,
		Kind:       string(s.Kind),
		StartTime:  strconv.FormatInt(s.StartTime.UnixMilli(), 10),
		EndTime:    strconv.FormatInt(s.EndTime.UnixMilli(), 10),
		Attributes: s.Attributes,
		Children:   encodeChildren(s.Children),
	}
}

func encodeChildren(children []*Span) []encodedSpan {
	res := make([]encodedSpan, len(children))
	for i, c := range children {
		res[i] = encodeSpan(*c)
	}
	return res
}

func (s *Span) UnmarshalJSON(data []byte) error {
	aux := encodedSpan{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("unmarshal span: %w", err)
	}

	return s.decodeSpan(aux)
}

func (s *Span) decodeSpan(aux encodedSpan) error {
	sid := trace.SpanID{}
	if aux.isValidID() {
		var err error
		sid, err = trace.SpanIDFromHex(aux.ID)
		if err != nil {
			return fmt.Errorf("unmarshal span: %w", err)
		}
	}

	children, err := decodeChildren(s, aux.Children, getCache())
	if err != nil {
		return fmt.Errorf("unmarshal span: %w", err)
	}

	startTime, err := getTimeFromString(aux.StartTime)
	if err != nil {
		return fmt.Errorf("unmarshal span: %w", err)
	}

	endTime, err := getTimeFromString(aux.EndTime)
	if err != nil {
		return fmt.Errorf("unmarshal span: %w", err)
	}

	s.ID = sid
	s.Name = aux.Name
	s.Kind = SpanKind(aux.Kind)
	s.StartTime = startTime.UTC()
	s.EndTime = endTime.UTC()
	s.Attributes = aux.Attributes
	s.Children = children

	return nil
}

func getTimeFromString(value string) (time.Time, error) {
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		// Maybe it is in RFC3339 format. Convert it for compatibility sake
		output, err := time.Parse(time.RFC3339, value)
		if err != nil {
			return time.Time{}, fmt.Errorf("could not convert string (%s) to time: %w", value, err)
		}

		return output, nil
	}

	return timing.ParseUnix(int64(parsedValue)), nil
}

func decodeChildren(parent *Span, children []encodedSpan, cache spanCache) ([]*Span, error) {
	if len(children) == 0 {
		return nil, nil
	}
	res := make([]*Span, len(children))
	for i, c := range children {
		if span, ok := cache.Get(c.ID); ok && span != nil {
			res[i] = span
			continue
		}

		span := &Span{
			Parent: parent,
		}
		if err := span.decodeSpan(c); err != nil {
			return nil, fmt.Errorf("unmarshal children: %w", err)
		}

		children, err := decodeChildren(span, c.Children, cache)
		if err != nil {
			return nil, fmt.Errorf("unmarshal children: %w", err)
		}

		span.Children = children
		res[i] = span

		cache.Set(span.ID.String(), span)
	}
	return res, nil
}

func (span Span) setMetadataAttributes() Span {
	if span.Attributes.values == nil {
		span.Attributes = NewAttributes()
	}

	span.Attributes.Set(QualitytraceMetadataFieldName, span.Name)
	span.Attributes.Set(QualitytraceMetadataFieldType, spanType(span.Attributes))
	span.Attributes.Set(QualitytraceMetadataFieldDuration, spanDuration(span))
	span.Attributes.Set(QualitytraceMetadataFieldStartTime, strconv.FormatInt(span.StartTime.UTC().UnixNano(), 10))
	span.Attributes.Set(QualitytraceMetadataFieldEndTime, strconv.FormatInt(span.EndTime.UTC().UnixNano(), 10))

	if span.Status != nil {
		span.Attributes.Set(QualitytraceMetadataFieldStatusCode, span.Status.Code)

		if span.Status.Description != "" {
			span.Attributes.Set(QualitytraceMetadataFieldStatusDescription, span.Status.Description)
		}
	}

	return span
}

func (span Span) setTriggerResultAttributes(result trigger.TriggerResult) Span {
	switch result.Type {
	case trigger.TriggerTypeHTTP:
		resp := result.HTTP
		jsonheaders, _ := json.Marshal(resp.Headers)
		span.Attributes.Set("qualitytrace.response.status", strconv.Itoa(resp.StatusCode))
		span.Attributes.Set("qualitytrace.response.body", resp.Body)
		span.Attributes.Set("qualitytrace.response.headers", string(jsonheaders))
	case trigger.TriggerTypeGRPC:
		resp := result.GRPC
		jsonheaders, _ := json.Marshal(resp.Metadata)
		span.Attributes.Set("qualitytrace.response.status", strconv.Itoa(resp.StatusCode))
		span.Attributes.Set("qualitytrace.response.body", resp.Body)
		span.Attributes.Set("qualitytrace.response.headers", string(jsonheaders))
	}

	return span
}
