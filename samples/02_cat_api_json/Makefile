.DEFAULT_GOAL := build
BIN_FILE=cats
build:
	go build -o dist/"${BIN_FILE}"
test:
	go test -v ./...
cover:
	go test -coverprofile cp.out
	go tool cover -html=cp.out
run:
	./"${BIN_FILE}"
clean:
	go clean
