test:
	go test ./... -v

deps:
	go mod tidy

vet:
	go vet ./...

lint:
	golangci-lint run
