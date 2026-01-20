# Automated Testing Go Demo [![CI](https://github.com/TallantM/go-framework-demo/actions/workflows/ci.yml/badge.svg)](https://github.com/TallantM/go-framework-demo/actions/workflows/ci.yml)

A demonstration of unit and integration testing in Go, using public APIs like jsonplaceholder.typicode.com for real-world scenarios, reusable utilities, and a GitHub Actions CI workflow for scalable quality assurance.

## Prerequisites
- Go 1.21+
- Git
- Visual Studio Code with Go extension
- Docker (for CI and optional local containerized testing)

Verify prerequisites:
- Go: `go version`
- Git: `git --version`
- Docker: `docker --version` (ensure running with `docker info`)

## Setup
1. Clone the repo: `git clone https://github.com/TallantM/go-framework-demo.git`
2. Navigate to root: `cd go-framework-demo`
3. Download dependencies: `go mod tidy`
4. Build the project: `go build` (optional, verifies compilation)

## Folder Structure
go-framework-demo/
├── internal/
│   └── utils/
│       └── helpers.go      # Reusable utilities (e.g., API helpers)
├── tests/
│   └── unit/
│       └── helpers_test.go # Unit tests
│   └── integration/
│       └── api_test.go     # Integration tests
├── .gitignore              # Ignores Go artifacts
├── Dockerfile              # Containerized build for testing
├── go.mod                  # Go module file
├── go.sum                  # Dependency checksums
├── README.md               # Documentation
└── .github/
    └── workflows/
        └── ci.yml          # GitHub Actions CI workflow

## Testing Overview
This repository demonstrates unit and integration testing in Go, using a public API (jsonplaceholder.typicode.com) for integration scenarios.
- **Unit Tests**: Isolated tests for utilities using Go's `testing` package and Testify, including table-driven patterns for extensive coverage and error handling (e.g., `tests/unit/helpers_test.go`).
- **Integration Tests**: HTTP requests to public APIs for real-world validation, with mocking for isolation and multiple scenarios (e.g., `tests/integration/api_test.go`).

## Running Tests Locally
For rapid development, run tests directly on your host:
go test ./...


For consistency with CI, use Docker (recommended before pushing):
docker build -t go-framework-demo .
docker run go-framework-demo

Run with race detection:
go test -race ./...

Generate coverage report:
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out


## CI/CD
GitHub Actions workflow in `.github/workflows/ci.yml` builds the Docker image and runs tests exclusively in the container on push/pull requests, ensuring environmental consistency.

## Troubleshooting
- **Test Failures**: Ensure dependencies are installed (`go mod tidy`). For integration tests, check network connectivity to jsonplaceholder.typicode.com.
- **CI Failures**: Review logs for module cache or Docker issues.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing
Contributions are welcome. Fork the repository, create a feature branch, and submit a pull request.