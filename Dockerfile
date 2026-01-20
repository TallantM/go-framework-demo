FROM docker:dind

# Install Go and other dependencies
RUN apk add --no-cache bash curl git
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin latest
RUN apk add --no-cache go

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Add a custom entrypoint script to wait for Docker daemon
RUN echo '#!/bin/bash' > /entrypoint.sh && \
    echo 'set -e' >> /entrypoint.sh && \
    echo 'dockerd --host=unix:///var/run/docker.sock --host=tcp://0.0.0.0:2375 > /dev/null 2>&1 &' >> /entrypoint.sh && \
    echo 'until docker info > /dev/null 2>&1; do sleep 1; done' >> /entrypoint.sh && \
    echo 'exec go test ./... -bench=. -short' >> /entrypoint.sh && \
    chmod +x /entrypoint.sh

ENTRYPOINT ["go", "test", "./...", "-coverprofile=coverage.out"]