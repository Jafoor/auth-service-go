package handlers

import (
	"auth-service/app/domain/user"
	"auth-service/config"
)

type Handlers struct {
	conf        *config.Config
	userService user.UserRepo
}

func NewHandler(
	conf *config.Config,
	userService user.UserRepo,
) *Handlers {
	return &Handlers{
		conf:        conf,
		userService: userService,
	}
}
