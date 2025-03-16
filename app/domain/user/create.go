package user

import (
	"auth-service/types"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (svc *service) Create(ctx context.Context, user types.SignUpUserPayload) error {
	isUserExists, _ := svc.userRepo.GetUserByEmail(ctx, user.Email)

	if isUserExists != nil {
		return fmt.Errorf("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	user.Password = string(hashedPassword)

	err = svc.userRepo.Create(ctx, user)

	if err != nil {
		return err
	}
	return nil
}
