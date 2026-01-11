# Technology Stack

## Language & Runtime

- **Go**: Version 1.12+ (as specified in go.mod and Dockerfile)
- **Standard Library**: Extensive use of net/http, encoding/json, context

## Dependencies

- **Minimal external dependencies**: The project uses only Go standard library
- **Module system**: Uses Go modules (go.mod) for dependency management

## Build System & Tools

### Development Commands

```bash
# Run tests
go test -v ./...

# Run tests with vendor mode
go test -mod vendor -v ./...

# Run linting
golint ./...

# Build the library
go build -mod vendor
```

### Docker Build

```bash
# Build Docker image (includes linting and testing)
docker build .
```

### CI/CD

- **GitHub Actions**: Automated builds on push/PR to master branch
- **Docker-based CI**: Uses multi-stage Docker build for testing and linting

## Code Quality Tools

- **golint**: Static analysis and style checking
- **go test**: Built-in testing framework with comprehensive test coverage
- **go vet**: Static analysis (implied in standard Go toolchain)

## Architecture Patterns

- **Service-oriented client**: Separate services for different API endpoints (PageService, ComponentService)
- **Shared HTTP client**: Common HTTP client with OAuth authentication
- **Context-aware**: All API calls accept context.Context for cancellation/timeouts
- **Pointer-based optionals**: Uses pointers for optional fields to distinguish between zero values and nil
