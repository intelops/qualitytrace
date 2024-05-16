FROM alpine

WORKDIR /app

COPY ./qualitytrace-server /app/qualitytrace-server

# COPY ./web/build ./html

EXPOSE 11633/tcp

ENTRYPOINT ["/app/qualitytrace-server", "serve"]
