package app

import (
	"github.com/intelops/qualitytrace/agent/tracedb"
	"github.com/intelops/qualitytrace/server/config"
	"github.com/intelops/qualitytrace/server/pkg/pipeline"
	"github.com/intelops/qualitytrace/server/testconnection"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

func buildDataStoreTestPipeline(
	driverFactory pipeline.DriverFactory[testconnection.Job],
	dsTestListener *testconnection.Listener,
	tracer trace.Tracer,
	newTraceDBFn tracedb.FactoryFunc,
	appConfig *config.AppConfig,
	meter metric.Meter,
) *testconnection.DataStoreTestPipeline {
	requestWorker := testconnection.NewDsTestConnectionRequest(tracer, newTraceDBFn, appConfig.DataStorePipelineTestConnectionEnabled())
	notifyWorker := testconnection.NewDsTestConnectionNotify(dsTestListener, tracer)

	pipeline := pipeline.New(testconnection.NewConfigurer(meter),
		pipeline.Step[testconnection.Job]{Processor: requestWorker, Driver: driverFactory.NewDriver("datastore_test_connection_request")},
		pipeline.Step[testconnection.Job]{Processor: notifyWorker, Driver: driverFactory.NewDriver("datastore_test_connection_notify")},
	)

	return testconnection.NewDataStoreTestPipeline(pipeline, dsTestListener)
}
