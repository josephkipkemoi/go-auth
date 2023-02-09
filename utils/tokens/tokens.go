package tokens

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)
var SECRETKEY = "maasai"

func GenerateToken() (string, error) {
	token := jwt.New(&jwt.SigningMethodHMAC{})
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "username"
	tokenString, err := token.SignedString(SECRETKEY)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken() bool {
	return false
}

func ExtractToken() string {
	return ""
}

