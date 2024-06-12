clean:
	find -type f -name 'Mock*.go' -print -delete
	go clean

.PHONY: mocks
mocks:
	mockery

code-generation:
	find -type f -name '*_impl.go' -print -delete
	go generate ./internal/adapters/*

test:
	go clean -testcache
	go test ./...

all: clean code-generation mocks
	go clean -testcache
	go test ./...

build:
	go build cmd/infofile/server.go