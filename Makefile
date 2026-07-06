.PHONY: build test test-v clean lint format format-check pre-commit-install ci install-cli pwa

BIN := bin/cais

test:
	go test ./... -race -count=1

test-v:
	go test ./... -v -count=1

build:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o $(BIN) ./cmd/cais

lint:
	golangci-lint run ./...

format:
	npm run format

format-check:
	npm run format:check

pre-commit-install:
	pre-commit install

ci: test lint format-check

install-cli:
	go install ./cmd/cais

pwa:
	go run ./cmd/pwagen

clean:
	rm -rf bin/ tmp/