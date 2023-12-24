BINARY_NAME=aio

build:
	@go build -o bin/$(BINARY_NAME) -v

run:
	@go run main.go

serve: generate
	@go run main.go serve

clean:
	@rm -rf bin/$(BINARY_NAME)

test:
	@go test -v ./...

coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
	@rm coverage.out

generate:
	@go generate ./...
	templ generate

air:
	@air -c .air.toml

.PHONY: build run clean test coverage generate air