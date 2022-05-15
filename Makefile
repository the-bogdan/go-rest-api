all: help

.PHONY: help
help:
	@echo "|--!!! Before other commands you need to execute \`make toolset\`!!!--|"
	@echo "toolset                  - install the necessary tools (test, lint, fmt)"
	@echo "\n|---- Code quality ----|"
	@echo "lint                     - run linters (golangci-lint)"
	@echo "fmt                      - code formatting (gofumpt)"
	@echo "test                     - run unit tests (gotestsum)"
	@echo "\n|---- Docs ----------|"
	@echo "swagger           		- update swagger docs"

#----------Code quality----------------------------------------------------------------------------#
.PHONY: toolset
toolset:
	@cd tools && go mod tidy && go generate tools.go

.PHONY: lint
lint:
	@cd app && ../tools/bin/golangci-lint run

.PHONY: fmt
fmt:
	@cd app && ../tools/bin/gofumpt -l -w .

.PHONY: test
test:
	@cd app && ../tools/bin/gotestsum ./...

#----------Docs------------------------------------------------------------------------------------#
.PHONY: swagger
swagger:
	@tools/bin/swag init -g ./app/cmd/server/main.go -o ./app/docs

#----------Build-----------------------------------------------------------------------------------#
.PHONY: build
build:
	@cd app && go build -o bin/app ./cmd/server

.PHONY: build-linux
build-linux:
	@cd app && GOOS=linux GOARCH=amd64 go build -o bin/app-linux ./cmd/server


