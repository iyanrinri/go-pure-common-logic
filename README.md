# Go Common Logic

A clean Go application implementing common algorithmic problems with proper software development practices.

## Project Structure

```
go-common-logic/
├── cmd/                    # Application entry points
│   └── main.go            # Main application
├── internal/              # Private application code
│   ├── handlers/          # HTTP request handlers
│   ├── services/          # Business logic layer
│   ├── models/            # Data structures
│   └── routes/            # Route configurations
├── tests/                 # Test files
├── go.mod                 # Go module definition
└── README.md             # Project documentation
```

## Running the Application

### Prerequisites
- Go 1.21 or higher

### Installation
```bash
git clone <repository-url>
cd go-common-logic
go mod tidy
```

### Run Server
```bash
go run cmd/main.go
```

The server will start on port 8081.

### Run Tests
```bash
go test ./tests/...
```

