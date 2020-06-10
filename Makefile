BINARY_NAME    = api
API_DIR     := ./api

.SILENT:
.DEFAULT_GOAL := help

.PHONY: help
help:
	$(info Available Commands:)
	$(info -> setup                   install dependencies modules)
	$(info -> format                  format go files)
	$(info -> build                   build binary)
	$(info -> test                    executes tests)
	$(info -> run                     starts server)
	$(info -> docker                  starts server on a docker image)

.PHONY: setup
install:
	go get -d -v ./...
	go install -v ./...
	go mod tidy -v

.PHONY: build
build:
	go build -v -o $(BINARY_NAME) $(API_DIR)
	chmod +x $(BINARY_NAME)

.PHONY: test
test:
	go test ./... -v -covermode=atomic

.PHONY: run
run:
	go run $(API_DIR)

.PHONY: docker
docker:
	docker build --build-arg root_dir=$(API_DIR) -t go-transactions .
	docker run --publish 8080:8080 --name go-transactions --rm go-transactions

.PHONY: format
format:
	go fmt ./...

# ignore unknown commands
%:
    @:
