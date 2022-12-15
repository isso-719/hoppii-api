.PHONY: test
test:
	go test -v ./...

.PHONY: test-without-user
test-without-user:
	go test `go list ./... | grep -v github.com/isso-719/hoppii-api/pkg/user`

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run ./...