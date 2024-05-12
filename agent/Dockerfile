ARG TRACETEST_VERSION=latest

FROM intelops/qualityTrace:${TRACETEST_VERSION}
WORKDIR /app

ENV TRACETEST_API_KEY ""

ENTRYPOINT [ "tracetest", "start", "--api-key", "$TRACETEST_API_KEY" ]

