ARG TRACETEST_VERSION=latest

FROM intelops/qualityTrace:${TRACETEST_VERSION}
WORKDIR /app

ENV TRACETEST_API_KEY ""

ENTRYPOINT [ "qualityTrace", "start", "--api-key", "$TRACETEST_API_KEY" ]

