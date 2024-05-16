FROM alpine

WORKDIR /app

COPY ./qualitytrace /app/qualitytrace

# Adding /app folder on $PATH to allow users to call qualitytrace cli on docker
ENV PATH="$PATH:/app"

ENTRYPOINT ["/app/qualitytrace"]
