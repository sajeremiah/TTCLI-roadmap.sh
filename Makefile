include .env

.PHONY: 

PROJECT_ROOT := $(PWD)

%:
	@:

build:
	@go build -o $(PROJECT_ROOT)/bin/task-cli $(PROJECT_ROOT)/cmd/ttcli/main.go

run: build
	@$(PROJECT_ROOT)/bin/task-cli $(filter-out $@,$(MAKECMDGOALS))

race:
	@go run -race $(PROJECT_ROOT)/cmd/ttcli/main.go