ARG TRACETEST_VERSION=latest

FROM intelops/qualitytrace:${TRACETEST_VERSION}
WORKDIR /app

ENV TRACETEST_API_KEY ""

ENTRYPOINT [ "qualitytrace", "start", "--api-key", "$TRACETEST_API_KEY" ]

