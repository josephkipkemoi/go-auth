package tokens

import (
	"go-auth/go-auth-api/models"
)

type Token struct {
	User *models.User
	JwtCode string
}

func (t Token) GenerateToken() string {
	return ""
} 