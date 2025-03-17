package user

import (
	"auth-service/app/adapter/rest/middlewares"
	"auth-service/types"
	"context"
)

type Service interface {
	Create(ctx context.Context, user types.SignUpUserPayload) error
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
	GetUserById(ctx context.Context, id int) (*types.User, error)
	LoginUser(ctx context.Context, user types.SignInUserPayload) (string, string, error)
	GetProfile(ctx context.Context, id int) (*types.ProfileResponse, error)
	ValidateToken(ctx context.Context, token string) (middlewares.AuthClaims, error)
	GenerateToken(user *types.User, expIn int, tokenType string) (string, error)
}
type UserRepo interface {
	Create(ctx context.Context, user types.SignUpUserPayload) error
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
	GetUserById(ctx context.Context, id int) (*types.User, error)
}
