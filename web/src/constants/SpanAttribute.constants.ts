import {SemanticAttributes, SemanticResourceAttributes} from '@opentelemetry/semantic-conventions';

export const TraceTestAttributes = {
  NAME: 'name',
  KIND: 'qualityTrace.span.kind',
  TRACETEST_SPAN_TYPE: 'qualityTrace.span.type',
  TRACETEST_SPAN_DURATION: 'qualityTrace.span.duration',
  TRACETEST_RESPONSE_STATUS: 'qualityTrace.response.status',
  TRACETEST_RESPONSE_BODY: 'qualityTrace.response.body',
  TRACETEST_RESPONSE_HEADERS: 'qualityTrace.response.headers',
  TRACETEST_SELECTED_SPANS_COUNT: 'qualityTrace.selected_spans.count',
};

export const Attributes: Record<string, string> = {
  ...SemanticAttributes,
  ...SemanticResourceAttributes,
  ...TraceTestAttributes,
  HTTP_REQUEST_HEADER: 'http.request.header.',
  HTTP_RESPONSE_HEADER: 'http.response.header',
};

export * from '@opentelemetry/semantic-conventions';
