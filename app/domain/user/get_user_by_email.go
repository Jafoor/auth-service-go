package user

import (
	"auth-service/types"
	"context"
)

func (svc *service) GetUserByEmail(ctx context.Context, email string) (*types.SignUpUser, error) {
	return svc.userRepo.GetUserByEmail(ctx, email)
}
