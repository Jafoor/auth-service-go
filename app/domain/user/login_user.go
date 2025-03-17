package user

import (
	"auth-service/types"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (svc *service) LoginUser(ctx context.Context, user types.SignInUserPayload) (string, string, error) {
	userInfo, err := svc.userRepo.GetUserByEmail(ctx, user.Email)

	if err != nil {
		return "", "", fmt.Errorf("failed to get user by email: %v", err)
	}

	if userInfo == nil {
		return "", "", fmt.Errorf("invalid password or email")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(user.Password)); err != nil {
		return "", "", fmt.Errorf("invalid password or email")
	}

	accessToken, err := svc.GenerateToken(userInfo, svc.JWT.AccessExpIn, "access")

	if err != nil {
		return "", "", fmt.Errorf("failed to generate token: %v", err)
	}

	refreshToken, err := svc.GenerateToken(userInfo, svc.JWT.RefreshExpIn, "refresh")

	if err != nil {
		return "", "", fmt.Errorf("failed to generate token: %v", err)
	}

	return accessToken, refreshToken, nil
}
