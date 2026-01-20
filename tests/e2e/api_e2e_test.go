package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/TallantM/go-framework-demo/internal/utils"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestE2EAPIWorkflow(t *testing.T) {
	ctx := context.Background()

	// Use WireMock container for custom API mocking
	req := testcontainers.ContainerRequest{
		Image:        "wiremock/wiremock:latest",
		ExposedPorts: []string{"8080/tcp"},
		WaitingFor:   wait.ForHTTP("/__admin/health").WithPort("8080/tcp"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	defer container.Terminate(ctx)

	// Get the host and mapped port
	host, err := container.Host(ctx)
	require.NoError(t, err)
	port, err := container.MappedPort(ctx, "8080")
	require.NoError(t, err)
	adminURL := fmt.Sprintf("http://%s:%s/__admin/mappings", host, port.Port())
	mockURL := fmt.Sprintf("http://%s:%s/mock-endpoint", host, port.Port())

	// Set up mock for POST
	postMapping := map[string]interface{}{
		"request": map[string]string{
			"method": "POST",
			"url":    "/mock-endpoint",
		},
		"response": map[string]interface{}{
			"status": 201,
			"body":   `{"id": 101, "title": "title", "body": "body"}`,
			"headers": map[string]string{
				"Content-Type": "application/json",
			},
		},
	}
	postJSON, _ := json.Marshal(postMapping)
	_, err = http.Post(adminURL, "application/json", bytes.NewBuffer(postJSON))
	require.NoError(t, err)

	// Set up mock for GET
	getMapping := map[string]interface{}{
		"request": map[string]string{
			"method": "GET",
			"url":    "/mock-endpoint",
		},
		"response": map[string]interface{}{
			"status": 200,
			"body":   `{"id": 101, "userId": 1, "title": "title", "body": "body"}`,
			"headers": map[string]string{
				"Content-Type": "application/json",
			},
		},
	}
	getJSON, _ := json.Marshal(getMapping)
	_, err = http.Post(adminURL, "application/json", bytes.NewBuffer(getJSON))
	require.NoError(t, err)

	// Set up mock for PATCH
	patchMapping := map[string]interface{}{
		"request": map[string]string{
			"method": "PATCH",
			"url":    "/mock-endpoint",
		},
		"response": map[string]interface{}{
			"status": 200,
			"body":   `{"id": 101, "title": "updated", "body": "body"}`,
			"headers": map[string]string{
				"Content-Type": "application/json",
			},
		},
	}
	patchJSON, _ := json.Marshal(patchMapping)
	_, err = http.Post(adminURL, "application/json", bytes.NewBuffer(patchJSON))
	require.NoError(t, err)

	// E2E workflow: POST, GET, PATCH with chained assertions
	postResp, err := utils.PostData(mockURL, "title", "body")
	require.NoError(t, err)
	require.Equal(t, 101, postResp.ID)

	getResp, err := utils.GetData(mockURL)
	require.NoError(t, err)
	require.Equal(t, postResp.Title, getResp.Title)

	updates := map[string]string{"title": "updated"}
	patchResp, err := utils.PatchData(mockURL, updates)
	require.NoError(t, err)
	require.Equal(t, "updated", patchResp.Title)
}
