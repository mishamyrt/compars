test:
	richgo test ./... -v
lint:
	golangci-lint run
coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
