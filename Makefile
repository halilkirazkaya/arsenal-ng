.PHONY: build

build:
	go build -ldflags "-s -w" -o bin/arsenal-ng ./cmd/arsenal-ng