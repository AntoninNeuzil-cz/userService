package main_test

import (
	"bytes"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	// Set up environment variables
	err := os.Setenv("DB_HOST", "localhost")
	if err != nil {
		assert.NoError(t, err, "Failed to set DB_HOST environment variable")
	}
	err = os.Setenv("DB_PORT", "8080")
	if err != nil {
		assert.NoError(t, err, "Failed to set DB_PORT environment variable")
	}

	// Start the application
	cmd := exec.Command("../user-service") // Path to your compiled binary
	cmd.Env = os.Environ()
	err = cmd.Start()
	assert.NoError(t, err, "Failed to start the application")

	// Wait for the application to start
	time.Sleep(2 * time.Second)

	// Test the Save endpoint
	savePayload := `{
		"external_id": "123e4567-e89b-12d3-a456-426614174000",
		"name": "John Doe",
		"email": "john.doe@example.com",
		"date_of_birth": "1990-01-01T00:00:00Z"
	}`
	resp, err := http.Post("http://localhost:8080/save", "application/json", bytes.NewBuffer([]byte(savePayload)))
	assert.NoError(t, err, "Failed to call Save endpoint")
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Test the Get endpoint
	getResp, err := http.Get("http://localhost:8080/123e4567-e89b-12d3-a456-426614174000")
	assert.NoError(t, err, "Failed to call Get endpoint")
	assert.Equal(t, http.StatusOK, getResp.StatusCode)

	// Stop the application
	err = cmd.Process.Kill()
	assert.NoError(t, err, "Failed to stop the application")
}
