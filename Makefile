test:
	@go test -v -race ./...

gen:
	@go generate ./...

clean:
	@go clean ./...
	@go mod tidy -v

lint:
	@go vet ./...
