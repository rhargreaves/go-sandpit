package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getResponse(t *testing.T) *http.Response {
	resp, err := http.Get("http://app:8080")
	require.NoError(t, err, "Failed to fetch URL")
	return resp
}

func TestServerIsRunning(t *testing.T) {
	resp := getResponse(t)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
}

func TestTextContentReturned(t *testing.T) {
	resp := getResponse(t)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Failed to read response body")

	assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal(t, "Hello, Docker with Go!\n", string(body))
}

func TestCustomHeaderReturned(t *testing.T) {
	resp := getResponse(t)
	defer resp.Body.Close()

	assert.Equal(t, "blah", resp.Header.Get("Another-Header"))
}

func TestJsonContentReturned(t *testing.T) {
	resp, err := http.Get("http://app:8080/json")
	require.NoError(t, err, "Failed to fetch URL")
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Failed to read response body")

	assert.Equal(t, "application/json; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.JSONEq(t, `{"message":"Hello, Docker with Go!"}`, string(body))
}
