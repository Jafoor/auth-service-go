package user

import (
	"auth-service/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *service) GenerateToken(user *types.User, expIn int, tokenType string) (string, error) {
	claims := jwt.MapClaims{
		"id":        user.Id,
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"type":      tokenType,
		"exp":       time.Now().Add(time.Duration(expIn) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.JWT.Secret))
}
