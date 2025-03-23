.DEFAULT_GOAL := run
.PHONY: run build

run: build
	./bin/api

build:
	go build -o ./bin/api ./cmd/api