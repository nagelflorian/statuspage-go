# Project Structure

## Root Level Files

- `statuspage.go` - Main client implementation with HTTP handling and authentication
- `component.go` - Component service for managing status page components
- `page.go` - Page service for managing status pages
- `strings.go` - Utility functions for string representation of structs
- `timestamp.go` - Custom timestamp type with JSON marshaling support

## Test Files

- `*_test.go` - Comprehensive test coverage for each service
- `*_property_test.go` - Property-based tests for correctness validation
- Test files follow Go naming convention with `_test.go` suffix
- Uses table-driven tests and mock HTTP servers for API testing
- Property-based tests validate universal correctness properties

## Configuration Files

- `go.mod` - Go module definition with minimal dependencies (requires Go 1.25+)
- `Dockerfile` - Multi-stage build for testing, linting, and building (uses golang:1.25 base)
- `.github/workflows/main.yml` - CI/CD pipeline configuration

## Code Organization Patterns

### Service Pattern

Each API resource has its own service file:

- `ComponentService` in `component.go`
- `PageService` in `page.go`
- Services are attached to the main `Client` struct

### Struct Definitions

- API response structs defined in their respective service files
- Request parameter structs follow `Update*Params` naming convention
- Request body structs follow `Update*RequestBody` naming convention

### Testing Structure

- Mock HTTP server setup in `statuspage_test.go`
- Helper functions for common test operations (`testMethod`, `testJSONMarshal`)
- Pointer helper functions (`String`, `Bool`, `Int32`, etc.) for test data creation
- Property-based tests for build system compatibility, API compatibility, and coverage preservation
- Sample client application for API backward compatibility validation

### Utility Files

- `strings.go` - Reflection-based string representation
- `timestamp.go` - Custom time handling for API compatibility

## Naming Conventions

- Services: `*Service` (e.g., `ComponentService`)
- API structs: Match API resource names (e.g., `Component`, `Page`)
- Update parameters: `Update*Params`
- Request bodies: `Update*RequestBody`
- Test helpers: Lowercase with descriptive names
