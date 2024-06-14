package app

import (
	"github.com/intelops/qualitytrace/agent/tracedb"
	"github.com/intelops/qualitytrace/server/config"
	"github.com/intelops/qualitytrace/server/datastore"
	"github.com/intelops/qualitytrace/server/executor"
	"github.com/intelops/qualitytrace/server/executor/pollingprofile"
	"github.com/intelops/qualitytrace/server/executor/testrunner"
	"github.com/intelops/qualitytrace/server/executor/tracepollerworker"
	"github.com/intelops/qualitytrace/server/executor/trigger"
	"github.com/intelops/qualitytrace/server/linter/analyzer"
	"github.com/intelops/qualitytrace/server/model"
	"github.com/intelops/qualitytrace/server/pkg/pipeline"
	"github.com/intelops/qualitytrace/server/subscription"
	"github.com/intelops/qualitytrace/server/test"
	"github.com/intelops/qualitytrace/server/testconnection"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

func buildTestPipeline(
	driverFactory pipeline.DriverFactory[executor.Job],
	pool *pgxpool.Pool,
	ppRepo *pollingprofile.Repository,
	dsRepo *datastore.Repository,
	lintRepo *analyzer.Repository,
	trRepo *testrunner.Repository,
	treRepo model.TestRunEventRepository,
	testRepo test.Repository,
	runRepo test.RunRepository,
	tracer trace.Tracer,
	subscriptionManager subscription.Manager,
	triggerRegistry *trigger.Registry,
	tracedbFactory tracedb.FactoryFunc,
	dataStoreTestPipeline *testconnection.DataStoreTestPipeline,
	appConfig *config.AppConfig,
	meter metric.Meter,
) *executor.TestPipeline {
	eventEmitter := executor.NewEventEmitter(treRepo, subscriptionManager)

	execTestUpdater := (executor.CompositeUpdater{}).
		Add(executor.NewDBUpdater(runRepo)).
		Add(executor.NewSubscriptionUpdater(subscriptionManager))

	workerMetricMiddlewareBuilder := executor.NewWorkerMetricMiddlewareBuilder(meter)

	assertionRunner := executor.NewAssertionRunner(
		execTestUpdater,
		executor.NewAssertionExecutor(tracer),
		executor.InstrumentedOutputProcessor(tracer),
		subscriptionManager,
		eventEmitter,
	)

	linterRunner := executor.NewLinterRunner(
		execTestUpdater,
		subscriptionManager,
		eventEmitter,
		lintRepo,
	)

	tracePollerStarterWorker := tracepollerworker.NewStarterWorker(
		eventEmitter,
		tracedbFactory,
		dsRepo,
		execTestUpdater,
		subscriptionManager,
		tracer,
		dataStoreTestPipeline,
	)

	traceFetcherWorker := tracepollerworker.NewFetcherWorker(
		eventEmitter,
		tracedbFactory,
		dsRepo,
		execTestUpdater,
		subscriptionManager,
		tracer,
		appConfig.TestPipelineTraceFetchingEnabled(),
	)

	tracePollerEvaluatorWorker := tracepollerworker.NewEvaluatorWorker(
		eventEmitter,
		tracedbFactory,
		dsRepo,
		execTestUpdater,
		subscriptionManager,
		tracepollerworker.NewSelectorBasedPollingStopStrategy(eventEmitter, tracepollerworker.NewSpanCountPollingStopStrategy()),
		tracer,
	)

	triggerResolverWorker := executor.NewTriggerResolverWorker(
		triggerRegistry,
		execTestUpdater,
		tracer,
		tracedbFactory,
		dsRepo,
		eventEmitter,
	)

	triggerExecuterWorker := executor.NewTriggerExecuterWorker(
		triggerRegistry,
		execTestUpdater,
		tracer,
		eventEmitter,
		appConfig.TestPipelineTriggerExecutionEnabled(),
	)

	triggerResultProcessorWorker := executor.NewTriggerResultProcessorWorker(
		tracer,
		subscriptionManager,
		eventEmitter,
		execTestUpdater,
	)

	cancelRunHandlerFn := executor.HandleRunCancelation(execTestUpdater, tracer, eventEmitter)

	queueBuilder := executor.NewQueueConfigurer().
		WithCancelRunHandlerFn(cancelRunHandlerFn).
		WithSubscriptor(subscriptionManager).
		WithDataStoreGetter(dsRepo).
		WithPollingProfileGetter(ppRepo).
		WithTestGetter(testRepo).
		WithRunGetter(runRepo).
		WithInstanceID(instanceID).
		WithMetricMeter(meter)

	pipeline := pipeline.New(queueBuilder,
		pipeline.Step[executor.Job]{Processor: workerMetricMiddlewareBuilder.New("trigger_resolver", triggerResolverWorker), Driver: driverFactory.NewDriver("trigger_resolve")},
		pipeline.Step[executor.Job]{Processor: workerMetricMiddlewareBuilder.New("trigger_executer", triggerExecuterWorker), Driver: driverFactory.NewDriver("trigger_execute")},
		pipeline.Step[executor.Job]{Processor: workerMetricMiddlewareBuilder.New("trigger_result_processor", triggerResultProcessorWorker), Driver: driverFactory.NewDriver("trigger_result")},
		pipeline.Step[executor.Job]{Processor: workerMetricMiddlewareBuilder.New("trace_poller_starter", tracePollerStarterWorker), Driver: driverFactory.NewDriver("tracePoller_start")},
		pipeline.Step[executor.Job]{Processor: workerMetricMiddlewareBuilder.New("trace_fetcher", traceFetcherWorker), Driver: driverFactory.NewDriver("tracePoller_fetch")},
		pipeline.Step[executor.Job]{Processor: workerMetricMiddlewareBuilder.New("trace_poller_evaluator", tracePollerEvaluatorWorker), Driver: driverFactory.NewDriver("tracePoller_evaluate"), InputQueueOffset: -1},
		pipeline.Step[executor.Job]{Processor: workerMetricMiddlewareBuilder.New("linter_runner", linterRunner), Driver: driverFactory.NewDriver("linterRunner")},
		pipeline.Step[executor.Job]{Processor: workerMetricMiddlewareBuilder.New("assertion_runner", assertionRunner), Driver: driverFactory.NewDriver("assertionRunner")},
	)

	const assertionRunnerStepIndex = 7

	return executor.NewTestPipeline(
		pipeline,
		subscriptionManager,
		pipeline.GetQueueForStep(assertionRunnerStepIndex), // assertion runner step
		runRepo,
		trRepo,
		ppRepo,
		dsRepo,
		cancelRunHandlerFn,
	)
}
