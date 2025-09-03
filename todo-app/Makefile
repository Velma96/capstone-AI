.PHONY: build-server build-cli run-server

build-server:
    go build -o bin/server ./cmd/server

build-cli:
    go build -o bin/todo ./cmd/cli

run-server:
    go run ./cmd/server

run-cli:
    go run ./cmd/cli $(filter-out $@,$(MAKECMDGOALS))