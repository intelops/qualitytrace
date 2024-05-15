FROM alpine

WORKDIR /app

COPY ./qualitytrace-server /app/qualitytrace-server
COPY ./qualitytrace /app/qualitytrace

COPY ./web/build ./html

# Adding /app folder on $PATH to allow users to call qualitytrace cli on docker
ENV PATH="$PATH:/app"

EXPOSE 11633/tcp

ENTRYPOINT ["/app/qualitytrace-server", "serve"]
