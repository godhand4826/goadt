.PHONY: test
test:
	go test -v -cover ./...

cover.html:
	go test -coverprofile cover.html ./...
	go tool cover -html=cover.html -o cover.html

doc: bin/godoc
	timeout 3s bin/godoc -http=localhost:8080&
	sleep 1
	-wget -q -r -np -N -E -p -k -e robots=off http://localhost:8080/pkg/goadt/
	mv localhost:8080 doc

.PHONY: lint
lint: bin/golangci-lint
	bin/golangci-lint run -v

bin/godoc:
	@echo "Installing godoc to ./bin"
	env GOBIN=$(CURDIR)/bin go install golang.org/x/tools/cmd/godoc@v0.23.0

bin/golangci-lint:
	@echo "Installing golangci-lint to ./bin"
	env GOBIN=$(CURDIR)/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.1
