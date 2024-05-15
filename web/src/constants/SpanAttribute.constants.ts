import {SemanticAttributes, SemanticResourceAttributes} from '@opentelemetry/semantic-conventions';

export const TraceTestAttributes = {
  NAME: 'name',
  KIND: 'qualitytrace.span.kind',
  TRACETEST_SPAN_TYPE: 'qualitytrace.span.type',
  TRACETEST_SPAN_DURATION: 'qualitytrace.span.duration',
  TRACETEST_RESPONSE_STATUS: 'qualitytrace.response.status',
  TRACETEST_RESPONSE_BODY: 'qualitytrace.response.body',
  TRACETEST_RESPONSE_HEADERS: 'qualitytrace.response.headers',
  TRACETEST_SELECTED_SPANS_COUNT: 'qualitytrace.selected_spans.count',
};

export const Attributes: Record<string, string> = {
  ...SemanticAttributes,
  ...SemanticResourceAttributes,
  ...TraceTestAttributes,
  HTTP_REQUEST_HEADER: 'http.request.header.',
  HTTP_RESPONSE_HEADER: 'http.response.header',
};

export * from '@opentelemetry/semantic-conventions';
