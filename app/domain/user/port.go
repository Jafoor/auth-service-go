package user

import (
	"auth-service/types"
	"context"
)

type Service interface {
	Create(ctx context.Context, user types.SignUpUserPayload) error
	GetUserByEmail(ctx context.Context, email string) (*types.SignUpUser, error)
}
type UserRepo interface {
	Create(ctx context.Context, user types.SignUpUserPayload) error
	GetUserByEmail(ctx context.Context, email string) (*types.SignUpUser, error)
}
