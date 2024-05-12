package poller

import (
	"context"

	"github.com/intelops/qualityTrace/agent/collector"
	"github.com/intelops/qualityTrace/agent/tracedb"
	"github.com/intelops/qualityTrace/agent/tracedb/connection"
	"github.com/intelops/qualityTrace/server/pkg/id"
	"github.com/intelops/qualityTrace/server/traces"
	"go.opentelemetry.io/otel/trace"
)

func NewInMemoryDatastore(cache collector.TraceCache) tracedb.TraceDB {
	return &inmemoryDatastore{cache}
}

type inmemoryDatastore struct {
	cache collector.TraceCache
}

// Close implements tracedb.TraceDB.
func (d *inmemoryDatastore) Close() error {
	return nil
}

// Connect implements tracedb.TraceDB.
func (d *inmemoryDatastore) Connect(ctx context.Context) error {
	return nil
}

// GetEndpoints implements tracedb.TraceDB.
func (d *inmemoryDatastore) GetEndpoints() string {
	return ""
}

// GetTraceByID implements tracedb.TraceDB.
func (d *inmemoryDatastore) GetTraceByID(ctx context.Context, traceID string) (traces.Trace, error) {
	spans, found := d.cache.Get(traceID)
	if !found || !d.cache.Exists(traceID) {
		return traces.Trace{}, connection.ErrTraceNotFound
	}

	return traces.FromSpanList(spans), nil
}

// GetTraceID implements tracedb.TraceDB.
func (d *inmemoryDatastore) GetTraceID() trace.TraceID {
	return id.NewRandGenerator().TraceID()
}

// Ready implements tracedb.TraceDB.
func (d *inmemoryDatastore) Ready() bool {
	return true
}

// ShouldRetry implements tracedb.TraceDB.
func (d *inmemoryDatastore) ShouldRetry() bool {
	return true
}
