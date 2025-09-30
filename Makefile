# Project variables
BIN_PLOTTER := bin/plotter
BIN_FARMER  := bin/farmer

# Default Go settings
GOFILES := $(shell find . -name '*.go' -not -path "./vendor/*")
GOMOD   := $(shell test -f go.mod && echo 1 || echo 0)

.PHONY: all build run-plotter run-farmer clean test deps fmt vet tidy

all: build

## Build both binaries
build: deps
	@mkdir -p bin
	go build -o $(BIN_PLOTTER) ./cmd/plotter
	go build -o $(BIN_FARMER)  ./cmd/farmer
	@echo "Built: $(BIN_PLOTTER) and $(BIN_FARMER)"

## Run plotter with defaults or overridden flags
run-plotter:
	go run ./cmd/plotter -size 10 -out plots

## Run farmer with a sample challenge
run-farmer:
	go run ./cmd/farmer -challenge test-challenge -plots plots

## Install/verify dependencies
deps:
ifeq ($(GOMOD),1)
	@echo "Using existing go.mod"
else
	@echo "Initializing go.mod"
	go mod init proof-of-space-poc
endif
	@echo "Ensuring modules..."
	go mod tidy

## Basic quality checks
fmt:
	gofmt -s -w $(GOFILES)

vet:
	go vet ./...

tidy:
	go mod tidy

## Run tests (placeholder; add tests under ./... as they are created)
test:
	go test ./...

## Clean build artifacts
clean:
	rm -rf bin

