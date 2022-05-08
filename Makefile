all: help

.PHONY: help
help:
	@echo "|---- Code quality ----|"
	@echo "toolset                  - устанавливает необходимые утилиты для работы с кодом (test, lint, fmt)"
	@echo "                         !!! Внимание, перед запуск следующих команд необходимо выполнить toolset "
	@echo "lint                     - запуск линтера (golangci-lint)"
	@echo "fmt                      - форматирование кода (gofumpt)"
	@echo "test                     - запуск тестов (gotestsum)"
	@echo "\n|---- Swagger -------|"
	@echo "swagger           		- Обновить документацию в swagger"

#----------Code quality----------------------------------------------------------------------------#
.PHONY: toolset
toolset:
	@cd tools && go mod tidy && go generate tools.go

.PHONY: lint
lint:
	cd app && @../tools/bin/golangci-lint run

.PHONY: fmt
fmt:
	@tools/bin/gofumpt -l -w .

.PHONY: test
test:
	@tools/bin/gotestsum ./...


#----------Swagger---------------------------------------------------------------------------------#

.PHONY: swagger
swagger:
	@tools/bin/swag init -g ./app/cmd/server/main.go -o ./app/docs

