package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (m *Manager) Use(middleware ...Middleware) *Manager {
	m.globalMiddlewares = append(m.globalMiddlewares, middleware...)
	return m
}

func (m *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	var h http.Handler = handler

	for _, middleware := range middlewares {
		h = middleware(h)
	}

	for _, middleware := range m.globalMiddlewares {
		h = middleware(h)
	}

	return h
}
