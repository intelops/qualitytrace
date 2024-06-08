FROM alpine

RUN apk --update add bash jq curl

WORKDIR /app
COPY ./dist/qualitytrace /app/qualitytrace
COPY ./testing/server-qualitytraceing ./qualitytraceing

WORKDIR /app/qualitytraceing
CMD ["/bin/sh", "/app/qualitytraceing/run.bash"]
