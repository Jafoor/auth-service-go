package rest

import (
	"auth-service/app/adapter/rest/handlers"
	"auth-service/app/adapter/rest/middlewares"
	"auth-service/config"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
)

type Server struct {
	Handlers *handlers.Handlers
	conf     *config.Config
	Wg       sync.WaitGroup
}

func NewServer(conf *config.Config, handlers *handlers.Handlers) *Server {
	return &Server{
		conf:     conf,
		Handlers: handlers,
	}
}

func (s *Server) Start() {
	manager := middlewares.NewManager()

	slog.Info("Rest Server Started")

	manager.Use(
		middlewares.Recover,
		middlewares.Logger,
	)

	mux := http.NewServeMux()

	initRoutes(mux, manager, s)

	handler := middlewares.EnableCors(mux)

	s.Wg.Add(1)

	go func() {
		defer s.Wg.Done()

		addr := fmt.Sprintf(":%d", s.conf.HttpPort)
		slog.Info(fmt.Sprintf("Listening at %s", addr))

		if err := http.ListenAndServe(addr, handler); err != nil {
			slog.Error(err.Error())
		}
	}()
}
