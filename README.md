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

## Features

### Two Sum Algorithm
Solves the classic two sum problem: given an array of integers and a target value, returns the indices of two numbers that add up to the target.

**Endpoint:** `POST /api/v1/two-sum`

**Time Complexity:** O(n)  
**Space Complexity:** O(n)

## API Usage

### Two Sum Request
```json
{
    "nums": [2, 7, 11, 15],
    "target": 9
}
```

### Two Sum Response
```json
{
    "indices": [0, 1]
}
```

### Error Response
```json
{
    "error": "no solution found"
}
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

## Example Usage

```bash
curl -X POST http://localhost:8081/api/v1/two-sum \
  -H "Content-Type: application/json" \
  -d '{"nums": [2, 7, 11, 15], "target": 9}'
```

Response:
```json
{
    "indices": [0, 1]
}
```

## Algorithm Details

The Two Sum implementation uses a hash map approach:

1. Iterate through the array once
2. For each number, calculate the complement (target - current number)
3. Check if complement exists in hash map
4. If found, return the indices
5. If not found, store current number and index in hash map
6. Continue until solution is found

This approach ensures optimal time complexity while maintaining code readability.

## Code Quality

The project follows Go best practices:
- Clean architecture with separated concerns
- Comprehensive test coverage
- Proper error handling
- Clear naming conventions
- Structured logging
- Input validation