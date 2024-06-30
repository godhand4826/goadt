.PHONY: test
test:
	go test -v -cover ./...

cover.html:
	go test -coverprofile cover.html
	go tool cover -html=cover.html -o cover.html

.PHONY: lint
lint: bin/golangci-lint
	bin/golangci-lint run -v

bin/golangci-lint:
	@echo "Installing golangci-lint to ./bin"
	env GOBIN=$(CURDIR)/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.1