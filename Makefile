.PHONY: init
init:
	go mod tidy

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: format
format:
	gofmt -w .
