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
	@echo "\n|---- Build ---------|"
	@echo "build           			- build executable binary for current system"
	@echo "build-linux           	- build executable binary for linux"
	@echo "\n|---- Docker --------|"
	@echo "docker-build           	- build docker image named go-rest-api"
	@echo "docker-rm            	- rm docker container named go-rest-api (even if it's running)"
	@echo "docker-run            	- run docker container named go-rest-api (if it's running then first remove old)"
	@echo "docker-login            	- login to docker registry to (you need change account name to your own)"
	@echo "docker-push            	- push docker image to registry"

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

#----------Docker----------------------------------------------------------------------------------#
.PHONY: docker-build
docker-build:
	@docker build --tag thebogdanp/go-rest-api .

.PHONY: docker-rm
docker-rm:
	@docker rm go-rest-api -f

.PHONY: docker-run
docker-run: docker-rm
	@docker run --name go-rest-api \
		--publish 127.0.0.1:9990:9990 \
		--env PORT=9990 \
		--env IS_PROD=false \
		--env LOG_LEVEL=info \
		--env HOST=0.0.0.0 \
		thebogdanp/go-rest-api

.PHONY: docker-login
docker-login:
	@docker login --username thebogdanp

.PHONY: docker-push
docker-push:
	@docker push thebogdanp/go-rest-api
