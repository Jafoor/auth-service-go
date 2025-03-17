package user

import (
	"auth-service/types"
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (svc *service) LoginUser(ctx context.Context, user types.SignInUserPayload) (string, error) {
	userInfo, err := svc.userRepo.GetUserByEmail(ctx, user.Email)

	if err != nil {
		return "", fmt.Errorf("failed to get user by email: %v", err)
	}

	if userInfo == nil {
		return "", fmt.Errorf("invalid password or email")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(user.Password)); err != nil {
		return "", fmt.Errorf("invalid password or email")
	}

	token, err := svc.generateToken(userInfo)

	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

func (s *service) generateToken(user *types.SignUpUser) (string, error) {
	claims := jwt.MapClaims{
		"id":        user.Id,
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"exp":       time.Now().Add(time.Duration(s.JWT.ExpIn) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.JWT.Secret))
}
