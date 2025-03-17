package user

import "auth-service/config"

type service struct {
	userRepo UserRepo
	JWT      config.Jwt
}

func NewService(userRepo UserRepo, jwt config.Jwt) Service {
	return &service{
		userRepo: userRepo,
		JWT:      jwt,
	}
}
