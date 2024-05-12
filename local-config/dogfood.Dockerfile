FROM alpine

RUN apk --update add bash jq curl

WORKDIR /app
COPY ./dist/qualityTrace /app/qualityTrace
COPY ./testing/server-qualityTraceing ./qualityTraceing

WORKDIR /app/qualityTraceing
CMD ["/bin/sh", "/app/qualityTraceing/run.bash"]
