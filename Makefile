BINARY_NAME    = bin
SERVER_DIR     := ./main

.SILENT:
.DEFAULT_GOAL := help

.PHONY: help
help:
	$(info Available Commands:)
	$(info -> setup                 install dependencies modules)
	$(info -> build                   build binary)
	$(info -> test                    run tests)
	$(info -> run                     run api)
	$(info -> format                  format go files)

.PHONY: setup
install:
	go mod tidy -v

.PHONY: build
build:
	go build -v $(ARGS) -o $(BINARY_NAME) $(SERVER_DIR)
	chmod +x $(BINARY_NAME)

.PHONY: test
test:
	go test ./... -v -race -covermode=atomic

.PHONY: run
run:
	go run $(SERVER_DIR)

.PHONY: format
format:
	go fmt ./...

# ignore unknown commands
%:
    @:
