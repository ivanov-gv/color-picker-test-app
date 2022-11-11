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

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: test100
test100:
	go test -v -count=100 ./...

.PHONY: race
race:
	go test -v -race -count=1 ./...

.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

.PHONY: gen
gen:
	mockgen -source=internal/pkg/repository/order/repository.go \
	-destination=internal/pkg/repository/order/mocks/mock_repository.go