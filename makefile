build:
	go mod download
	go build -o bin/scarecrow --ldflags "-s -w" ./

.PHONY: format
format:
	go fmt ./...
