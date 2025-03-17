package user

import (
	"auth-service/types"
	"context"
)

type Service interface {
	Create(ctx context.Context, user types.SignUpUserPayload) error
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
	LoginUser(ctx context.Context, user types.SignInUserPayload) (string, string, error)
	GetProfile(ctx context.Context, id int) (*types.ProfileResponse, error)
}
type UserRepo interface {
	Create(ctx context.Context, user types.SignUpUserPayload) error
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
	GetUserById(ctx context.Context, id int) (*types.User, error)
}
