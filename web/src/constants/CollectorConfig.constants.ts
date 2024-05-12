import {SupportedDataStores} from 'types/DataStore.types';

export const qualityTrace = `# OTLP for Tracetest
  otlp/qualityTrace:
    endpoint: qualityTrace:4317 # Send traces to Tracetest. Read more in docs here:  https://docs.qualityTrace.io/configuration/connecting-to-data-stores/opentelemetry-collector
    tls:
      insecure: true`;

export const Lightstep = (traceTestBlock: string) => `receivers:
  otlp:
    protocols:
      grpc:
      http:

processors:
  batch:
    timeout: 100ms

exporters:
  logging:
    logLevel: debug

  ${traceTestBlock}

  # OTLP for Lightstep
  otlp/lightstep:
    endpoint: ingest.lightstep.com:443
    headers:
      "lightstep-access-token": "<lightstep_access_token>" # Send traces to Lightstep. Read more in docs here: https://docs.lightstep.com/otel/otel-quick-start

service:
  pipelines:
    traces/qualityTrace:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp/qualityTrace]
    traces/lightstep:
      receivers: [otlp]
      processors: [batch]
      exporters: [logging, otlp/lightstep]
`;

export const OtelCollector = (traceTestBlock: string) => `receivers:
  otlp:
    protocols:
      grpc:
      http:

processors:
  batch:
    timeout: 100ms

exporters:
  ${traceTestBlock}

service:
  pipelines:
    traces/1:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp/qualityTrace]
`;

export const NewRelic = (traceTestBlock: string) => `receivers:
  otlp:
    protocols:
      grpc:
      http:

processors:
  batch:
    timeout: 100ms

exporters:
  logging:
    logLevel: debug

  ${traceTestBlock}

  # OTLP for New Relic
  otlp/newrelic:
    endpoint: otlp.nr-data.net:443
    headers:
      api-key: <new_relic_ingest_licence_key> # Send traces to New Relic.
      # Read more in docs here: https://docs.newrelic.com/docs/more-integrations/open-source-telemetry-integrations/opentelemetry/opentelemetry-setup/#collector-export
      # And here: https://docs.newrelic.com/docs/more-integrations/open-source-telemetry-integrations/opentelemetry/collector/opentelemetry-collector-basic/

service:
  pipelines:
    traces/qualityTrace:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp/qualityTrace]
    traces/newrelic:
      receivers: [otlp]
      processors: [batch]
      exporters: [logging, otlp/newrelic]
`;

export const Datadog = (traceTestBlock: string) => `receivers:
  otlp:
    protocols:
      http:
      grpc:

processors:
  batch:
    send_batch_max_size: 100
    send_batch_size: 10
    timeout: 10s

exporters:
  ${traceTestBlock}

  # Datadog exporter
  datadog:
    api:
      site: datadoghq.com
      key: <datadog_API_key> # Add here you API key for Datadog
      # Read more in docs here: https://docs.datadoghq.com/opentelemetry/otel_collector_datadog_exporter
service:
  pipelines:
    traces/qualityTrace:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp/qualityTrace]
    traces/datadog:
      receivers: [otlp]
      processors: [batch]
      exporters: [datadog]
`;

export const Honeycomb = (traceTestBlock: string) => `receivers:
  otlp:
    protocols:
      grpc:
      http:

processors:
  batch:
    timeout: 100ms

exporters:
  logging:
    logLevel: debug

  ${traceTestBlock}

  # OTLP for Honeycomb
  otlp/honeycomb:
    endpoint: "api.honeycomb.io:443"
    headers:
      "x-honeycomb-team": "YOUR_API_KEY"
      # Read more in docs here: https://docs.honeycomb.io/getting-data-in/otel-collector/

service:
  pipelines:
    traces/qualityTrace:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp/qualityTrace]
    traces/honeycomb:
      receivers: [otlp]
      processors: [batch]
      exporters: [logging, otlp/honeycomb]
`;

export const AzureAppInsights = (traceTestBlock: string) => `receivers:
otlp:
  protocols:
    grpc:
    http:

processors:
  batch:

exporters:
  azuremonitor:
    instrumentation_key: <your-instrumentation-key>

  ${traceTestBlock}

service:
  pipelines:
    traces/qualityTrace:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp/qualityTrace]
    traces/appinsights:
      receivers: [otlp]
      exporters: [azuremonitor]
`;

export const Signoz = (traceTestBlock: string) => `receivers:
  otlp:
    protocols:
      grpc:
      http:

processors:
  batch:
    timeout: 100ms

exporters:
  logging:
    logLevel: debug

  ${traceTestBlock}

  # OTLP for Signoz
  otlp/signoz:
    endpoint: address-to-your-signoz-server:4317 # Send traces to Signoz. Read more in docs here: https://signoz.io/docs/tutorial/opentelemetry-binary-usage-in-virtual-machine/#opentelemetry-collector-configuration
    tls:
      insecure: true

service:
  pipelines:
    traces/qualityTrace:
      receivers: [otlp]
      processors: [batch]
      exporters: [logging, otlp/qualityTrace]
    traces/signoz:
      receivers: [otlp]
      processors: [batch]
      exporters: [logging, otlp/signoz]
`;

export const Dynatrace = (traceTestBlock: string) => `receivers:
  otlp:
    protocols:
      grpc:
      http:

processors:
  batch:
    timeout: 100ms

exporters:
  logging:
    verbosity: detailed

  ${traceTestBlock}

  # OTLP for Dynatrace
  otlphttp/dynatrace:
    endpoint: https://abc12345.live.dynatrace.com/api/v2/otlp # Send traces to Dynatrace. Read more in docs here: https://www.dynatrace.com/support/help/extend-dynatrace/opentelemetry/collector#configuration
    headers:
      Authorization: "Api-Token dt0c01.sample.secret"  # Requires "openTelemetryTrace.ingest" permission

service:
  pipelines:
    traces/qualityTrace:
      receivers: [otlp]
      processors: [batch]
      exporters: [logging, otlp/qualityTrace]
    traces/dynatrace:
      receivers: [otlp]
      processors: [batch]
      exporters: [logging, otlphttp/dynatrace]
`;

export const Instana = (traceTestBlock: string) => `receivers:
  otlp:
    protocols:
      grpc:
      http:

processors:
  batch:
    timeout: 100ms

exporters:
  logging:
    verbosity: detailed

  ${traceTestBlock}

  # OTLP for Instana
  # Send traces to Instana. Read more in docs here: https://www.ibm.com/docs/en/instana-observability/current?topic=opentelemetry-sending-data-instana-backend
  otlp/instana:
    endpoint: otlp-XXXX-saas.instana.io:4317 # it can be one of the SaaS environments. Look on the Instana doc link for more details.
    headers:
      x-instana-key: some-key # it is the Instana Agent Key provided by Instana UI

service:
  pipelines:
    traces/qualityTrace:
      receivers: [otlp]
      processors: [batch]
      exporters: [logging, otlp/qualityTrace]
    traces/instana:
      receivers: [otlp]
      processors: [batch]
      exporters: [logging, otlp/instana]
`;

export const CollectorConfigMap = {
  [SupportedDataStores.AzureAppInsights]: AzureAppInsights(qualityTrace),
  [SupportedDataStores.Datadog]: Datadog(qualityTrace),
  [SupportedDataStores.Dynatrace]: Dynatrace(qualityTrace),
  [SupportedDataStores.Honeycomb]: Honeycomb(qualityTrace),
  [SupportedDataStores.Instana]: Instana(qualityTrace),
  [SupportedDataStores.Lightstep]: Lightstep(qualityTrace),
  [SupportedDataStores.NewRelic]: NewRelic(qualityTrace),
  [SupportedDataStores.OtelCollector]: OtelCollector(qualityTrace),
  [SupportedDataStores.Signoz]: Signoz(qualityTrace),
} as const;

export const CollectorConfigFunctionMap = {
  [SupportedDataStores.AzureAppInsights]: AzureAppInsights,
  [SupportedDataStores.Datadog]: Datadog,
  [SupportedDataStores.Dynatrace]: Dynatrace,
  [SupportedDataStores.Honeycomb]: Honeycomb,
  [SupportedDataStores.Instana]: Instana,
  [SupportedDataStores.Lightstep]: Lightstep,
  [SupportedDataStores.NewRelic]: NewRelic,
  [SupportedDataStores.OtelCollector]: OtelCollector,
  [SupportedDataStores.Signoz]: Signoz,
} as const;
