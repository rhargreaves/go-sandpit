package main

import (
	"net/http"
	"testing"
)

func TestServerIsRunning(t *testing.T) {
	resp, err := http.Get("http://app:8080")
	if err != nil {
		t.Fatalf("Failed to reach server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}
