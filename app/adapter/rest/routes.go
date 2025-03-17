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

	mux.Handle(
		"POST /api/v1/signup",
		manager.With(
			http.HandlerFunc(s.Handlers.SignUpUser)),
	)

	mux.Handle(
		"POST /api/v1/signin",
		manager.With(
			http.HandlerFunc(s.Handlers.SignInUser)),
	)

	mux.Handle(
		"POST /api/v1/refresh",
		manager.With(
			http.HandlerFunc(s.Handlers.RefreshToken)),
	)

	mux.Handle(
		"GET /api/v1/profile",
		manager.With(
			http.HandlerFunc(s.Handlers.GetUserDetails),
			middlewares.Authenticate,
		),
	)
}
