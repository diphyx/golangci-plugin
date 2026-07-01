.PHONY: init
init:
	go mod tidy

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: format
format:
	gofmt -w .

.PHONY: test
test:
	go test ./...

.PHONY: verify
verify:
	sh verify.sh

.PHONY: publish
publish:
	sh publish.sh
