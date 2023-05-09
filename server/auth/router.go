package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/mnrva-dev/owltier.com/server/middleware"
)

func BuildRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/login", Login)
	r.Post("/register", Register)
	r.Group(func(r chi.Router) {
		r.Use(middleware.TokenValidater)
		r.Post("/delete", DeleteAccount)
	})
	r.Group(func(r chi.Router) {
		r.Use(middleware.RefreshValidator)
		r.Post("/token/refresh", Refresh)
	})
	r.Get("/token/validate", Validate)

	return r
}
