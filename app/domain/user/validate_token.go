package user

import (
	"auth-service/app/adapter/rest/middlewares"
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func (svc *service) ValidateToken(ctx context.Context, token string) (middlewares.AuthClaims, error) {
	var claims middlewares.AuthClaims
	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(svc.JWT.Secret), nil
	})

	if err != nil {
		return middlewares.AuthClaims{}, fmt.Errorf("invalid token: %v", err)
	}

	if !parsedToken.Valid {
		return middlewares.AuthClaims{}, fmt.Errorf("invalid token")
	}

	if _, ok := parsedToken.Claims.(*middlewares.AuthClaims); !ok {
		return middlewares.AuthClaims{}, fmt.Errorf("invalid token claims type")
	}

	return claims, nil
}
