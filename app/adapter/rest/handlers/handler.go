package handlers

import "auth-service/config"

type Handlers struct {
	conf *config.Config
}

func NewHandler(
	conf *config.Config,
) *Handlers {
	return &Handlers{
		conf: conf,
	}
}
