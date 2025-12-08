# Project: go-demo

## 1. Project Overview

This is a Go project (`github.com/yungyu16/go-demo`) that serves as a collection of various code examples and learning demonstrations for Go version 1.24.

The repository is structured into two main directories:
- `cmd/`: Contains small, runnable programs demonstrating specific commands or language features.
- `pkg/`: Traditionally for library code, but in this project, it also contains many runnable `main` packages, each illustrating a different concept (e.g., concurrency, HTTP servers, timers).

The primary purpose of this repository is for studying and experimenting with the Go language and its ecosystem. It is not a single, cohesive application but a set of independent examples.

## 2. Building and Running

There are no centralized build scripts or a main application. Each file containing a `main` package can be run individually.

### Running an Example

To run a specific example, use the `go run` command followed by the file path.

**Examples:**
```bash
# Run the echo web server example
go run cmd/echo.go

# Run the http server example from the misc package
go run pkg/misc/http.go
```

### Building an Example

Similarly, you can build a binary for any runnable example using `go build`.

```bash
# Build the echo example
go build cmd/echo.go
```

### Testing

Currently, there are no test files (`*_test.go`) in the project. If tests are added in the future, they can be run using the standard `go test` command.

```bash
# Run all tests (if any are added)
go test ./...
```

## 3. Development Conventions

- **Project Structure:** The code follows a standard Go project layout with `cmd` and `pkg` directories.
- **Runnable Packages:** A key convention in this repository is that many packages, including those in `pkg/misc`, are created as `package main`. This indicates that they are designed to be executed directly for demonstration purposes, rather than being imported as libraries.
- **Code Style:** The code appears to follow standard Go formatting, consistent with `gofmt`.
- **Dependencies:** Project dependencies are managed using Go Modules (`go.mod` and `go.sum`).
