FROM golang:1.19.3
COPY . /app
WORKDIR /app

RUN go env -w CGO_ENABLED=0 &&\
    go env -w GOOS=linux &&\
    go env -w GOARCH=amd64

ARG ENV=dev
COPY ./config/${ENV}.yaml ./config.yaml
ENV CONFIG_PATH="config.yaml"

CMD ["go", "test", "-v", "./integration-test/..."]
