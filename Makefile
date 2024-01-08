# Variables
BINARY_NAME := project-starter
CMD_PATH := .
PKG_PATH := ./pkg
TEST_PATH := ./tests
BUILD_PATH := project-starter
VERSION := 1.0.0

# Commands

run:
	go run -race .

build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) $(CMD_PATH)

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_PATH)

linux:
	env GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) $(CMD_PATH)

lint:
	gofmt -s -w .
# staticcheck ./...
	go test -race -vet=off ./...
	go vet ./...
	golangci-lint run --timeout 5m -v