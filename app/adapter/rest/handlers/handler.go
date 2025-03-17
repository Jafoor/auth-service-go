package handlers

import (
	"auth-service/app/domain/user"
	"auth-service/config"
)

type Handlers struct {
	conf        *config.Config
	userService user.Service
}

func NewHandler(
	conf *config.Config,
	userService user.Service,
) *Handlers {
	return &Handlers{
		conf:        conf,
		userService: userService,
	}
}
