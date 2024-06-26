package tracepollerworker

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/intelops/qualitytrace/agent/tracedb"
	"github.com/intelops/qualitytrace/agent/tracedb/connection"
	"github.com/intelops/qualitytrace/server/datastore"
	"github.com/intelops/qualitytrace/server/executor"
	"github.com/intelops/qualitytrace/server/model"
	"github.com/intelops/qualitytrace/server/pkg/pipeline"
	"github.com/intelops/qualitytrace/server/resourcemanager"
	"github.com/intelops/qualitytrace/server/subscription"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type workerState struct {
	eventEmitter        executor.EventEmitter
	newTraceDBFn        tracedb.FactoryFunc
	dsRepo              resourcemanager.Current[datastore.DataStore]
	updater             executor.RunUpdater
	subscriptionManager subscription.Manager
	tracer              trace.Tracer
	inputQueue          pipeline.Enqueuer[executor.Job]
}

func emitEvent(ctx context.Context, state *workerState, event model.TestRunEvent) {
	err := state.eventEmitter.Emit(ctx, event)
	if err != nil {
		log.Printf("[TracePoller] failed to emit %s event: error: %s", event.Type, err.Error())
	}
}

func getTraceDB(ctx context.Context, state *workerState) (tracedb.TraceDB, error) {
	ds, err := state.dsRepo.Current(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot get default datastore: %w", err)
	}

	tdb, err := state.newTraceDBFn(ds)
	if err != nil {
		return nil, fmt.Errorf(`cannot get tracedb from DataStore config with ID "%s": %w`, ds.ID, err)
	}

	return tdb, nil
}

func handleError(ctx context.Context, job executor.Job, err error, state *workerState, span trace.Span) {
	log.Printf("[TracePoller] Test %s Run %d, Error: %s", job.Test.ID, job.Run.ID, err.Error())

	span.RecordError(err)
	span.SetAttributes(attribute.String("qualitytrace.run.trace_poller.error", err.Error()))
}

func handleDBError(err error) {
	if err == nil {
		return
	}

	log.Printf("[TracePoller] DB error when polling traces: %s\n", err.Error())
}

func populateSpan(span trace.Span, job executor.Job, reason string, done *bool) {
	spanCount := 0
	if job.Run.Trace != nil {
		spanCount = len(job.Run.Trace.Flat)
	}

	attrs := []attribute.KeyValue{
		attribute.String("qualitytrace.run.trace_poller.trace_id", job.Run.TraceID.String()),
		attribute.String("qualitytrace.run.trace_poller.span_id", job.Run.SpanID.String()),
		attribute.String("qualitytrace.run.trace_poller.test_id", string(job.Test.ID)),
		attribute.Int("qualitytrace.run.trace_poller.amount_retrieved_spans", spanCount),
	}

	if done != nil {
		attrs = append(attrs, attribute.Bool("qualitytrace.run.trace_poller.succesful", *done))
	}

	if reason != "" {
		attrs = append(attrs, attribute.String("qualitytrace.run.trace_poller.finish_reason", reason))
	}

	span.SetAttributes(attrs...)
}

func isTraceNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, connection.ErrTraceNotFound)
}
