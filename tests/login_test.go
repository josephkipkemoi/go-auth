package tests

import (
	"bytes"
	"encoding/json"
	factory "go-auth/go-auth-api/database/factory" // load user factory
	"go-auth/go-auth-api/env"
	h "go-auth/go-auth-api/handlers/auth" // load auth handlers
	r "go-auth/go-auth-api/routes"        // load routes
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestUnauthenticatedUserCannotLogin tests that an unauthenticated user cannot log in
func TestUnauthenticatedUserCannotLogin(t *testing.T) {
	url := env.GetDevAppUrl()
	router := r.SetupRouter()
	w := httptest.NewRecorder()

	user := factory.MakeUser()
	
	u := &h.User{
		PhoneNumber: int(user.PhoneNumber),
		Password: "password",
	}

	d, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	body := bytes.NewReader(d)

	// // user request
	req := httptest.NewRequest("POST", url + "api/v1/login", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, "application/json", w.Header().Get("Content-Type"), "Should have application/json header set")
	assert.Equal(t, http.StatusBadRequest, w.Code, "Should return Bad Request (400): Invalid Credentials")
}	

// TestAuthenticatedUserCanLogIn tests a user with right credentials can login
func TestAuthenticatedUserCanLogIn(t *testing.T) {
	url := env.GetDevAppUrl()
	router := r.SetupRouter()
	w := httptest.NewRecorder()

	user := factory.MakeUser()
	
	u := &h.User{
		PhoneNumber: int(user.PhoneNumber),
		Password: string(user.Password),
	}

	d, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	body := bytes.NewReader(d)

	// // user request
	req := httptest.NewRequest("POST", url + "api/v1/login", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, "application/json", w.Header().Get("Content-Type"), "Should have application/json header set")
	assert.Equal(t, http.StatusNoContent, w.Code, "Should return success code (200): User found and Match records")
	assert.Equal(t, "bearer jwt_token", w.Header().Get("Authorization"), "Should have authorization header once user is logged in")
}