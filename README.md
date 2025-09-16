# gingerale

A minimal web server using the Gin framework for Go, providing simple HTTP routing with JSON and string responses.

## Features
- Lightweight HTTP server with Gin framework
- Health check endpoint for monitoring
- Simple routing examples
- Ready for development and deployment

## Requirements
- Go 1.20+ (recommended Go 1.23.5)
- Git (for cloning and version control)

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd gingerale
   ```

2. Initialize and download dependencies:
   ```bash
   go mod tidy
   ```

## Usage

### Development
Run the development server (local port 8080):
```bash
# Recommended to specify local directory for Go build cache in restricted environments
GOCACHE=$(pwd)/.gocache go run .
```

### Production Build
Build a production binary:
```bash
go build -o gingerale
./gingerale
```

## API Endpoints

### Root Route
- **GET** `/`
- **Description**: Welcome message
- **Response**: Plain text message

```bash
curl -i http://localhost:8080/
```
Response: `Welcome to gingerale with Gin!`

### Health Check
- **GET** `/ping`
- **Description**: Health check endpoint for monitoring
- **Response**: JSON status message

```bash
curl -i http://localhost:8080/ping
```
Response: `{"message":"pong"}`

## Project Structure
```
.
├── .gocache/          # Go build cache (generated)
├── .gitignore         # Git ignore rules
├── CLAUDE.md          # Claude Code guidance
├── LICENSE            # License file
├── README.md          # This file
├── go.mod             # Go module definition
├── go.sum             # Go module checksums
├── main.go            # Main application entry point
└── gingerale          # Compiled binary (generated)
```

## Development

### Testing
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

### Code Quality
```bash
# Format code
go fmt ./...

# Lint code
go vet ./...
```

## Deployment Considerations

For production environments, consider adding:
- Graceful shutdown handling
- Structured logging (logrus, zap)
- Configuration management (viper, environment variables)
- Middleware for CORS, authentication, rate limiting
- Docker containerization
- Health check improvements with database connectivity tests

## Notes
- If your environment has network access restrictions, initial `go get` or `go mod tidy` may require proxy configuration
- The server runs on `0.0.0.0:8080` by default
- Use environment variables like `PORT` to configure the listening port in production
