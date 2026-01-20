# Automated Testing Go Demo [![CI](https://github.com/TallantM/go-framework-demo/actions/workflows/ci.yml/badge.svg)](https://github.com/TallantM/go-framework-demo/actions/workflows/ci.yml)

A client project demonstration showcasing unit, integration, and e2e testing in Go, incorporating reusable utilities, layered test suites, and a GitHub Actions CI workflow with Docker containerization for reliable quality assurance.

## Prerequisites
- Go 1.21+ (verify with `go version`).
- Git (verify with `git --version`).
- Visual Studio Code with Go extension.
- Docker (verify with `docker --version`; required for CI and optional local containerized runs).

## Setup
1. Clone the repo: `git clone https://github.com/TallantM/go-framework-demo.git`
2. Navigate to root: `cd go-framework-demo`
3. Download dependencies: `go mod tidy`
4. Build the project: `go build` (optional, verifies compilation)

## Folder Structure
```text
ggo-framework-demo/
├── internal/
│   └── utils/
│       └── helpers.go      # Reusable utilities (e.g., API helpers)
├── tests/
│   └── unit/
│       └── helpers_test.go # Unit tests
│   └── integration/
│       └── api_test.go     # Integration tests
│   └── benchmarks/
│       └── helpers_bench_test.go # Benchmark tests
│   └── e2e/
│       └── api_e2e_test.go # End-to-end tests
│   └── fuzz/
│       └── helpers_fuzz.go # Fuzz tests
├── .gitignore              # Ignores Go artifacts
├── Dockerfile              # Containerized build for testing
├── go.mod                  # Go module file
├── go.sum                  # Dependency checksums
├── README.md               # Documentation
└── .github/
    └── workflows/
        └── ci.yml          # GitHub Actions CI workflow
```

## Testing Overview
This repository demonstrates unit and integration testing in Go, using public APIs like jsonplaceholder.typicode.com for real-world scenarios, reusable utilities, and a GitHub Actions CI workflow for scalable quality assurance.
- **Unit Tests**: Isolated tests for utilities using Go's `testing` package and Testify, including table-driven patterns for extensive coverage and error handling (e.g., `tests/unit/helpers_test.go`).
- **Integration Tests**: HTTP requests to public APIs for real-world validation, with mocking for isolation and multiple scenarios (e.g., `tests/integration/api_test.go`).
- **Benchmark Tests**: Performance measurements with parallelism and allocation reporting (e.g., `tests/benchmarks/helpers_bench_test.go`).
- **End-to-End (E2E) Tests**: Full workflow validations using Testcontainers-Go for mocked environments (e.g., `tests/e2e/api_e2e_test.go`).
- **Fuzz Tests**: Randomized input generation for robustness (e.g., `tests/fuzz/helpers_fuzz.go`).

## Running Tests Locally

For rapid development, run tests directly on your host:
```bash
go test ./...
```

For consistency with CI, use Docker (recommended before pushing):
```bash
docker build -t go-framework-demo .
docker run go-framework-demo
```

Run with race detection:
```bash
go test -race ./...
```

Generate coverage report:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

Run benchmarks:
```bash
go test -bench=. -benchmem ./tests/benchmarks
```

Run fuzz tests (run for a few seconds, then Ctrl+C to stop):
```bash
go test -fuzz=. ./tests/fuzz
```

## CI/CD
GitHub Actions workflow in `.github/workflows/ci.yml` builds the Docker image and runs tests exclusively in the container on push/pull requests, ensuring environmental consistency. It includes continuous performance monitoring with benchstat for benchmark comparisons.

## Troubleshooting
- **Test Failures**: Ensure dependencies are installed (`go mod tidy`). For integration tests, check network connectivity to jsonplaceholder.typicode.com. For E2E tests, verify Docker is running.
- **CI Failures**: Review logs for module cache or Docker issues, including benchmark comparisons or coverage uploads.

## Contributing
Contributions are welcome. Fork the repository, create a feature branch, and submit a pull request.
