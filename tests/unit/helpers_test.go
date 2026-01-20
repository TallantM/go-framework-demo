package unit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TallantM/go-framework-demo/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestAddTableDriven(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{name: "positive numbers", a: 2, b: 3, expected: 5},
		{name: "zero values", a: 0, b: 0, expected: 0},
		{name: "negative numbers", a: -1, b: 1, expected: 0},
		{name: "large numbers", a: 1000000, b: 1000000, expected: 2000000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := utils.Add(tc.a, tc.b)
			assert.Equal(t, tc.expected, result, "Addition failed for %s", tc.name)
		})
	}
}

func TestPostDataMocked(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": 101, "title": "mocktitle", "body": "mockbody"}`))
	}))
	defer mockServer.Close()

	response, err := utils.PostData(mockServer.URL, "title", "body")
	assert.NoError(t, err)
	assert.Equal(t, 101, response.ID)
}

func TestPostDataErrorHandling(t *testing.T) {
	_, err := utils.PostData("invalid-url", "title", "body")
	assert.Error(t, err, "Expected error for invalid URL")
}

func TestPostDataNonSuccessStatus(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(`{"error": "bad request"}`))
	}))
	defer mockServer.Close()

	_, err := utils.PostData(mockServer.URL, "title", "body")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "API request failed with status 400")
}

func TestPostDataDecodeError(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`)) // Malformed JSON
	}))
	defer mockServer.Close()

	_, err := utils.PostData(mockServer.URL, "title", "body")
	assert.Error(t, err)
}

func TestGetDataMocked(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": 1, "userId": 1, "title": "mocktitle", "body": "mockbody"}`))
	}))
	defer mockServer.Close()

	response, err := utils.GetData(mockServer.URL)
	assert.NoError(t, err)
	assert.Equal(t, 1, response.ID)
}

func TestGetDataErrorHandling(t *testing.T) {
	_, err := utils.GetData("invalid-url")
	assert.Error(t, err, "Expected error for invalid URL")
}

func TestGetDataNonSuccessStatus(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound) // 404
	}))
	defer mockServer.Close()

	_, err := utils.GetData(mockServer.URL)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "API request failed with status 404")
}

func TestGetDataDecodeError(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`)) // Malformed JSON
	}))
	defer mockServer.Close()

	_, err := utils.GetData(mockServer.URL)
	assert.Error(t, err)
}

func TestPatchDataMocked(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": 1, "title": "updatedtitle", "body": "updatedbody"}`))
	}))
	defer mockServer.Close()

	updates := map[string]string{"title": "updatedtitle", "body": "updatedbody"}
	response, err := utils.PatchData(mockServer.URL, updates)
	assert.NoError(t, err)
	assert.Equal(t, 1, response.ID)
}

func TestPatchDataErrorHandling(t *testing.T) {
	updates := map[string]string{"title": "updated"}
	_, err := utils.PatchData("invalid-url", updates)
	assert.Error(t, err, "Expected error for invalid URL")
}

func TestPatchDataNonSuccessStatus(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest) // 400
	}))
	defer mockServer.Close()

	updates := map[string]string{"title": "updated"}
	_, err := utils.PatchData(mockServer.URL, updates)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "API request failed with status 400")
}

func TestPatchDataDecodeError(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`)) // Malformed JSON
	}))
	defer mockServer.Close()

	updates := map[string]string{"title": "updated"}
	_, err := utils.PatchData(mockServer.URL, updates)
	assert.Error(t, err)
}
