package tests

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"

	"go-auth/go-auth-api/env"             // load environment variables
	h "go-auth/go-auth-api/controllers/auth" // load auth handlers
	r "go-auth/go-auth-api/routes"        // load routes
	"go-auth/go-auth-api/utils/faker"     // load faker

	"github.com/stretchr/testify/assert"
)

// Test can register new user
func TestNewUserRegistration(t *testing.T) {
	url := 	env.GetDevAppUrl()
	router := r.SetupRouter()

	w := httptest.NewRecorder()
	// Create user from User Struct
	i := &h.RegistrationInput{
		PhoneNumber: faker.PhoneNumber(),
		Password: "j",
	}
	
	// json encode
	e, err := json.Marshal(&i)
	if err != nil {
		fmt.Println(err)
	}
	// Return io.Reader to post to route
	bodyReader := bytes.NewReader(e)
	// Send body
	req, _ := http.NewRequest("POST", url + "api/v1/register", bodyReader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code, "Should return 201 status code after user is created")
	assert.Equal(t, "201 Created", w.Result().Status, "Should contain 'User Created' success/status message" )
	assert.Equal(t, "bearer jwt_token", w.Header().Get("Authorization"), "Should return populated Authorization header if authenticated")
}

