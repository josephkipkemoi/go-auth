package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	h "go-auth/go-auth-api/controllers"
	r "go-auth/go-auth-api/routes"
	"go-auth/go-auth-api/models"

	"github.com/stretchr/testify/assert"
)

// Test can load/view landing page with appropriate message
func TestLanding(t *testing.T) {
	models.ConnectDB()
	router := r.SetupRouter()
	
	w := httptest.NewRecorder()
	
	req, _ := http.NewRequest("GET", "/api/", nil)
	router.ServeHTTP(w, req)

	message := h.Message{Message: h.WelcomeMessage}

	body ,err := json.Marshal(message)

	if err != nil {
		panic("Error")
	}

	assert.Equal(t, 200, w.Code, "Should return HTTP success status code 200")
	assert.Equal(t, string(body), w.Body.String(), "Should return correct body string")
}