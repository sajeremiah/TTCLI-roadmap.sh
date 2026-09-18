include .env

.PHONY: 

PROJECT_ROOT := $(PWD)

build:
	@go build -o $(PROJECT_ROOT)/bin/ttcli $(PROJECT_ROOT)/cmd/ttcli/main.go

run: build
	@$(PROJECT_ROOT)/bin/ttcli

race:
	@go run -race $(PROJECT_ROOT)/cmd/ttcli/main.go