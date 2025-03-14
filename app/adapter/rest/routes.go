package rest

import (
	"auth-service/app/adapter/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middlewares.Manager, s *Server) {
	mux.Handle(
		"GET /hello",
		manager.With(
			http.HandlerFunc(s.Handlers.Hello),
		),
	)
}
