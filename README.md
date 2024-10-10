# Quizer

A simple quiz application with a REST API and CLI client.

## Prerequisites

- Go 1.23 or higher
- Make (optional, for using Makefile commands)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/popeskul/quizer.git
   cd quizer
   ```

2. Install dependencies:
   ```
   go mod download
   ```

## Running the Application

1. Start the server:
   ```
   go run cmd/quiz/main.go
   ```

2. Use the CLI client:
   ```
   go run cmd/cli/main.go get    # Get a quiz
   go run cmd/cli/main.go submit # Submit answers
   ```

## Testing

Run all tests:
```
go test ./...
```

## API Documentation

Access Swagger UI at `http://localhost:8080/swagger/` when the server is running.

## What's Tested

- Core domain logic
- Use cases
- API handlers
- In-memory repository
- CLI commands

## Dependencies

- github.com/go-chi/chi/v5: HTTP router
- github.com/spf13/cobra: CLI framework
- github.com/spf13/viper: Configuration management
- github.com/stretchr/testify: Testing toolkit
- go.uber.org/mock: Mocking framework for tests

## License

This project is licensed under the MIT License.