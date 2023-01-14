package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	r "go-auth/go-auth-api/routes"

	"github.com/stretchr/testify/assert"
)

// Test can load/view landing page with appropriate message
func TestLanding(t *testing.T) {
	router := r.SetupRouter()
	
	w := httptest.NewRecorder()
	
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)
	
	message := `{"message":"Golang Auth API"}`

	assert.Equal(t, 200, w.Code, "Should return HTTP success status code 200")
	assert.Equal(t, message, w.Body.String(), "Should return correct body string")
}