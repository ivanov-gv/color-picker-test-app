BIN_PATH?=./bin/app
ENV?=local

.PHONY: clean
clean:
	rm -f ${BIN_PATH}

.PHONY: build
build: clean
	go build -tags migrate -o ${BIN_PATH} ./cmd

.PHONY: run
run: build
	CONFIG_PATH="./config/${ENV}.yaml" ${BIN_PATH}

.PHONY: run-docker
run-docker:
	docker compose up