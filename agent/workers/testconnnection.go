package workers

import (
	"context"
	"fmt"
	"log"

	"github.com/intelops/qualitytrace/agent/telemetry"

	"github.com/intelops/qualitytrace/agent/client"
	"github.com/intelops/qualitytrace/agent/event"
	"github.com/intelops/qualitytrace/agent/proto"
	"github.com/intelops/qualitytrace/agent/tracedb"
	"github.com/intelops/qualitytrace/server/model"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type TestConnectionWorker struct {
	client   *client.Client
	logger   *zap.Logger
	observer event.Observer
	tracer   trace.Tracer
	meter    metric.Meter
}

type TestConnectionOption func(*TestConnectionWorker)

func WithTestConnectionLogger(logger *zap.Logger) TestConnectionOption {
	return func(w *TestConnectionWorker) {
		w.logger = logger
	}
}

func WithTestConnectionObserver(observer event.Observer) TestConnectionOption {
	return func(w *TestConnectionWorker) {
		w.observer = observer
	}
}

func WithTestConnectionTracer(tracer trace.Tracer) TestConnectionOption {
	return func(w *TestConnectionWorker) {
		w.tracer = tracer
	}
}

func WithTestConnectionMeter(meter metric.Meter) TestConnectionOption {
	return func(w *TestConnectionWorker) {
		w.meter = meter
	}
}

func NewTestConnectionWorker(client *client.Client, opts ...TestConnectionOption) *TestConnectionWorker {
	worker := &TestConnectionWorker{
		client:   client,
		tracer:   telemetry.GetNoopTracer(),
		logger:   zap.NewNop(),
		observer: event.NewNopObserver(),
		meter:    telemetry.GetNoopMeter(),
	}

	for _, opt := range opts {
		opt(worker)
	}

	return worker
}

func (w *TestConnectionWorker) Test(ctx context.Context, request *proto.DataStoreConnectionTestRequest) error {
	ctx, span := w.tracer.Start(ctx, "TestConnectionRequest Worker operation")
	defer span.End()

	runCounter, _ := w.meter.Int64Counter("qualitytrace.agent.testconnectionworker.runs")
	runCounter.Add(ctx, 1)

	errorCounter, _ := w.meter.Int64Counter("qualitytrace.agent.testconnectionworker.errors")

	w.logger.Debug("Received datastore connection test request")
	w.observer.StartDataStoreConnection(request)

	datastoreConfig, err := convertProtoToDataStore(request.Datastore)
	if err != nil {
		w.logger.Error("Invalid datastore", zap.Error(err))
		w.observer.EndDataStoreConnection(request, err)
		span.RecordError(err)
		errorCounter.Add(ctx, 1)

		return err
	}

	w.logger.Debug("Converted datastore", zap.Any("datastore", datastoreConfig))

	if datastoreConfig == nil {
		err = fmt.Errorf("invalid datastore: nil")

		w.logger.Error("nil datastore", zap.Error(err))
		w.observer.EndDataStoreConnection(request, err)
		span.RecordError(err)
		errorCounter.Add(ctx, 1)

		return err
	}

	dsFactory := tracedb.Factory(nil)
	ds, err := dsFactory(*datastoreConfig)
	if err != nil {
		w.logger.Error("Invalid datastore", zap.Error(err))
		log.Printf("Invalid datastore: %s", err.Error())
		w.observer.EndDataStoreConnection(request, err)
		span.RecordError(err)
		errorCounter.Add(ctx, 1)

		return err
	}
	w.logger.Debug("Created datastore", zap.Any("datastore", ds))

	response := &proto.DataStoreConnectionTestResponse{
		RequestID:  request.RequestID,
		Successful: false,
		Steps:      nil,
	}

	if testableTraceDB, ok := ds.(tracedb.TestableTraceDB); ok {
		w.logger.Debug("Datastore is testable")
		connectionResult := testableTraceDB.TestConnection(ctx)
		w.logger.Debug("Tested datastore", zap.Any("connectionResult", connectionResult))
		success, steps := convertConnectionResultToProto(connectionResult)
		w.logger.Debug("Converted connection result", zap.Bool("success", success), zap.Any("steps", steps))

		response = &proto.DataStoreConnectionTestResponse{
			RequestID:  request.RequestID,
			Successful: success,
			Steps:      steps,
		}
	} else {
		w.logger.Debug("Datastore is not testable")
	}

	w.logger.Debug("Sending datastore connection test result", zap.Any("response", response))
	err = w.client.SendDataStoreConnectionResult(ctx, response)
	if err != nil {
		w.logger.Error("Could not send datastore connection test result", zap.Error(err))
		w.observer.Error(err)
		span.RecordError(err)
		errorCounter.Add(ctx, 1)
	} else {
		w.logger.Debug("Sent datastore connection test result")
	}

	w.observer.EndDataStoreConnection(request, nil)
	return nil
}

func convertConnectionResultToProto(connectionResult model.ConnectionResult) (bool, *proto.DataStoreConnectionTestSteps) {
	steps := &proto.DataStoreConnectionTestSteps{
		PortCheck:      convertConnectionResultStepToProto(connectionResult.PortCheck),
		Connectivity:   convertConnectionResultStepToProto(connectionResult.Connectivity),
		Authentication: convertConnectionResultStepToProto(connectionResult.Authentication),
		FetchTraces:    convertConnectionResultStepToProto(connectionResult.FetchTraces),
	}

	return connectionResult.HasSucceed(), steps
}

func convertConnectionResultStepToProto(step model.ConnectionTestStep) *proto.DataStoreConnectionTestStep {
	errorMsg := ""
	if step.Error != nil {
		errorMsg = step.Error.Error()
	}
	return &proto.DataStoreConnectionTestStep{
		Passed:  step.Passed,
		Status:  string(step.Status),
		Message: step.Message,
		Error:   errorMsg,
	}
}
