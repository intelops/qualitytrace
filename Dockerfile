FROM alpine

WORKDIR /app

COPY ./qualityTrace-server /app/qualityTrace-server
COPY ./qualityTrace /app/qualityTrace

COPY ./web/build ./html

# Adding /app folder on $PATH to allow users to call qualityTrace cli on docker
ENV PATH="$PATH:/app"

EXPOSE 11633/tcp

ENTRYPOINT ["/app/qualityTrace-server", "serve"]
