FROM golang:1.19.3 AS build-env
WORKDIR /go/src/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux make build

FROM scratch
COPY --from=build-env /go/src/bin/app /app
COPY migrations migrations

ARG ENV=dev
COPY ./config/${ENV}.yaml /config.yaml
ENV CONFIG_PATH="config.yaml"

EXPOSE 8080
CMD ["/app"]