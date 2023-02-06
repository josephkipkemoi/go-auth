package tests

import (
	"os"
	"testing"

	_ "go-auth-api/go-auth/env"       // load environment variables

	"github.com/stretchr/testify/assert"
)

// Test environment variables are loaded correctly
func TestEnvVariables(t *testing.T) {
	devUrlPath := os.Getenv("GO_AUTH_API_DEV_URL")

	assert.Equal(t, "http://localhost:8080/", devUrlPath, "It should have the correct full URL path")
}