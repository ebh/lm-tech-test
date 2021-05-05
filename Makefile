.PHONY: fmt lint test test-mutation

test:
	go test -v -race

test-mutation:
	go get -t -v github.com/zimmski/go-mutesting/...
	go-mutesting ./...

lint:
	golangci-lint run

fmt:
	`go list -f {{.Target}} mvdan.cc/gofumpt` -w .
