package user

import (
	"auth-service/types"
	"context"
)

func (svc *service) GetUserById(ctx context.Context, id int) (*types.User, error) {
	return svc.userRepo.GetUserById(ctx, id)
}
