# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Building and Running
- Run the development server: `GOCACHE=$(pwd)/.gocache go run .`
- Build dependencies: `go mod tidy`
- Build binary: `go build -o gingerale`

### Testing
- Run tests: `go test ./...`
- Run tests with coverage: `go test -cover ./...`

### Code Quality
- Format code: `go fmt ./...`
- Lint code: `go vet ./...`

## Architecture Overview

This is a minimal Gin framework HTTP server application with the following structure:

### Core Components
- **main.go**: Single-file application containing the Gin router setup and route handlers
- **go.mod**: Go module definition with Gin framework dependency (v1.10.0)

### Routes
- `GET /`: Welcome message endpoint
- `GET /ping`: Health check endpoint returning JSON response

### Key Patterns
- Uses Gin's default router with built-in middleware
- Runs on port 8080 by default
- Simple inline handler functions for routes
- JSON responses for API endpoints, string responses for simple messages

### Dependencies
- Primary: `github.com/gin-gonic/gin` for HTTP routing and middleware
- Go version: 1.23.5

### Development Notes
- Uses local `.gocache` directory for Go build cache in restricted environments
- No additional configuration files or complex build setup
- Single binary deployment model