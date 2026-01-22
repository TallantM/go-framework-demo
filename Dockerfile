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
RUN echo '#!/bin/sh' > /tmp/entrypoint.sh && \
    echo 'set -e' >> /tmp/entrypoint.sh && \
    echo 'dockerd --host=unix:///var/run/docker.sock --host=tcp://0.0.0.0:2375 > /dev/null 2>&1 &' >> /tmp/entrypoint.sh && \
    echo 'until docker info > /dev/null 2>&1; do sleep 1; done' >> /tmp/entrypoint.sh && \
    echo 'export DOCKER_HOST=unix:///var/run/docker.sock' >> /tmp/entrypoint.sh && \
    echo 'export TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE=/var/run/docker.sock' >> /tmp/entrypoint.sh && \
    echo 'export TESTCONTAINERS_RYUK_DISABLED=true' >> /tmp/entrypoint.sh && \
    echo 'exec go test ./... -bench=. -coverprofile=coverage.out -coverpkg=github.com/TallantM/go-framework-demo/internal/utils' >> /tmp/entrypoint.sh && \
    chmod +x /tmp/entrypoint.sh && \
    cp /tmp/entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]