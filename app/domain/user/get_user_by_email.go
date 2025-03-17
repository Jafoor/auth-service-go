package user

import (
	"auth-service/types"
	"context"
)

func (svc *service) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	return svc.userRepo.GetUserByEmail(ctx, email)
}
