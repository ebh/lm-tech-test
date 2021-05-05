.PHONY: fmt lint test

test:
	go test -v -race

lint:
	golangci-lint run

fmt:
	`go list -f {{.Target}} mvdan.cc/gofumpt` -w .
