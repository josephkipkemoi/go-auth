package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go-auth/go-auth-api/env"
	_ "go-auth/go-auth-api/env"           // load environment variables
	h "go-auth/go-auth-api/handlers/auth" // load auth handlers
	r "go-auth/go-auth-api/routes"        // load routes

	"github.com/stretchr/testify/assert"
)

// Test can register new user
func TestNewUserRegistration(t *testing.T) {
	url := 	env.GetDevAppUrl()
	router := r.SetupRouter()

	w := httptest.NewRecorder()
	// Create user from User Struct
	u := h.User{
		Username: "Joseph",
		Email: "jkemboe@gmail.com",
		Password: "123",
		CreatedAt: time.Now(),
	}
	// json encode
	e, _ := json.Marshal(u)
	
	bodyReader := bytes.NewReader(e)
	 
	// Send body
	req, _ := http.NewRequest("POST", url + "api/v1/register", bodyReader)
	router.ServeHTTP(w, req)

	reqBody, _ := ioutil.ReadAll(req.Body)
	
	assert.Equal(t, e, reqBody, "Should have user registratoin fields")
	assert.Equal(t, 201, w.Code, "Should return 201 status code after user is created")
}