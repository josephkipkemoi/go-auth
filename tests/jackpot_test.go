package tests

import (
	"go-auth/go-auth-api/env"
	"go-auth/go-auth-api/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPostJackpotMarkets(t *testing.T) {
	url := env.GetDevAppUrl()
	r := routes.SetupRouter()
	w := httptest.NewRecorder()

	req := httptest.NewRequest("POST", url + "api/v1/jackpots", nil)
	r.ServeHTTP(w, req)

	assert.Contains(t, w.Header().Get("Content-Type"), "application/json", "Should have content-type data format set to json")
	assert.Contains(t, w.Header().Get("Authorization"), "", "Should have valid jwt authorization header")
	assert.JSONEq(t,`{"market":"Mega Jackpot"}`,w.Body.String(), "Should have JSON BODY")
	assert.Equal(t,http.StatusCreated, w.Code, "Should return resource created http status code")
}