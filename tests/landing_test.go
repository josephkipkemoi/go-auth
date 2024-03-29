package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	h "go-auth-api/go-auth/controllers"
	r "go-auth-api/go-auth/routes"

	"github.com/stretchr/testify/assert"
)

// Test can load/view landing page with appropriate message
func TestLanding(t *testing.T) {
	router := r.SetupRouter()
	
	w := httptest.NewRecorder()
	
	req, _ := http.NewRequest("GET", "/api/v1/", nil)
	router.ServeHTTP(w, req)

	message := h.Message{Message: h.WelcomeMessage}

	body ,err := json.Marshal(message)

	if err != nil {
		panic("Error")
	}

	assert.Equal(t, 200, w.Code, "Should return HTTP success status code 200")
	assert.Equal(t, string(body), w.Body.String(), "Should return correct body string")
}