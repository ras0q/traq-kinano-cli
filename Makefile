SHELL := /bin/bash

build:
	@go build -v ./...

.PHONY: go-gen
go-gen:
	@go generate ./...

.PHONY: ent-add
ent-add:
	@read -p "Schema name > " schema && \
	go run entgo.io/ent/cmd/ent@latest init --target ./ent/schema $$schema

.PHONY: ent-gen
ent-gen:
	@go run entgo.io/ent/cmd/ent@latest generate ./ent/schema
