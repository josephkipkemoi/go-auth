package landing_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	r "go-auth/go-auth-api/routes"
)

// Test can load/view landing page with appropriate message
func TestLanding(t *testing.T) {
	router := r.SetupRouter()
	
	w := httptest.NewRecorder()
	
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Should return HTTP success status code 200")
	assert.Equal(t, "Golang Auth API", w.Body.String(), "Should return correct body string")
}