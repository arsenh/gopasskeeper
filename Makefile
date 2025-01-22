# Makefile for Golang project

# Variables
APP_NAME := gopasskeeper
SRC_DIR := ./cmd/main.go


# Run the application
.PHONY: run
run:
	@./$(APP_NAME)

# Build the application
.PHONY: build
build:
	@go build -o $(APP_NAME) $(SRC_DIR)
	@echo "Build finished."

# Clean up build artifacts
.PHONY: clean
clean:
	@rm -f $(APP_NAME)
	@echo "Cleanup."

# Run tests
.PHONY: test
test:
	@go test ./...

# Format code
.PHONY: fmt
fmt:
	@go fmt ./...
