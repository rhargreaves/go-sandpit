package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServerIsRunning(t *testing.T) {
	resp, err := http.Get("http://app:8080")
	require.NoError(t, err, "Failed to reach server")
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
}
