.PHONY: build run test clean fmt vet

# Build the application
build:
	go build -o bin/app cmd/main.go

# Run the application
run:
	go run cmd/main.go

# Run tests
test:
	go test -v ./tests/...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./tests/...
	go tool cover -html=coverage.out -o coverage.html

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run all checks
check: fmt vet test

# Development setup
dev-setup: deps
	@echo "Development environment ready"