build:
	go mod download
	go test -v ./...
	go mod tidy