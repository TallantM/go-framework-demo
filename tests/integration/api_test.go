package integration

import (
	"testing"

	"github.com/TallantM/go-framework-demo/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestPostIntegration(t *testing.T) {
	response, err := utils.PostData("https://jsonplaceholder.typicode.com/posts", "title", "body")
	assert.NoError(t, err)
	assert.NotZero(t, response.ID)
	assert.Equal(t, "title", response.Title)
	assert.Equal(t, "body", response.Body)
}

func TestInvalidPostIntegration(t *testing.T) {
	// jsonplaceholder returns 404 for invalid endpoints, triggering an error in PostData
	_, err := utils.PostData("https://jsonplaceholder.typicode.com/invalid", "title", "body")
	assert.Error(t, err, "Expected error for invalid endpoint")
}

func TestGetIntegration(t *testing.T) {
	response, err := utils.GetData("https://jsonplaceholder.typicode.com/posts/1")
	assert.NoError(t, err)
	assert.Equal(t, 1, response.ID)
	assert.NotEmpty(t, response.Title)
	assert.NotEmpty(t, response.Body)
}

func TestInvalidGetIntegration(t *testing.T) {
	_, err := utils.GetData("https://jsonplaceholder.typicode.com/invalid")
	assert.Error(t, err, "Expected error for invalid endpoint")
}

func TestPatchIntegration(t *testing.T) {
	updates := map[string]string{"title": "updated title", "body": "updated body"}
	response, err := utils.PatchData("https://jsonplaceholder.typicode.com/posts/1", updates)
	assert.NoError(t, err)
	assert.Equal(t, 1, response.ID) // jsonplaceholder echoes updates in response
}

func TestInvalidPatchIntegration(t *testing.T) {
	updates := map[string]string{"title": "updated"}
	_, err := utils.PatchData("https://jsonplaceholder.typicode.com/invalid", updates)
	assert.Error(t, err, "Expected error for invalid endpoint")
}
