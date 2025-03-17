package middlewares

import (
	"auth-service/app/adapter/rest/utils"
	"auth-service/config"
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "user"

// AuthClaims represents the JWT claims structure.
type AuthClaims struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Type     string `json:"type"`
	FistName string `json:"firstName"`
	LastName string `json:"lastName"`
	jwt.RegisteredClaims
}

// unauthorized sends a standardized unauthorized response.
func unauthorized(w http.ResponseWriter) {
	utils.SendJSON(w, http.StatusUnauthorized, map[string]any{
		"success": false,
		"message": "Unauthorized",
	})
}

// extractToken extracts the JWT token from the request header or query parameter.
func extractToken(r *http.Request) string {
	// Check Authorization header first
	header := r.Header.Get("Authorization")
	if header != "" {
		tokens := strings.Split(header, " ")
		if len(tokens) == 2 && strings.ToLower(tokens[0]) == "bearer" {
			return tokens[1]
		}
	}

	// Fallback to query parameter
	return r.URL.Query().Get("token")
}

// Authenticate is a middleware to validate JWT tokens.
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conf := config.GetConfig()

		// Extract token from the request
		tokenStr := extractToken(r)
		if tokenStr == "" {
			unauthorized(w)
			return
		}

		// Parse and validate the token
		var claims AuthClaims
		token, err := jwt.ParseWithClaims(
			tokenStr,
			&claims,
			func(t *jwt.Token) (interface{}, error) {
				return []byte(conf.JWT.Secret), nil
			},
		)

		if err != nil || !token.Valid {
			unauthorized(w)
			return
		}

		// Add claims to the request context
		ctx := context.WithValue(r.Context(), UserKey, &claims)
		wrappedRequest := r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, wrappedRequest)
	})
}
