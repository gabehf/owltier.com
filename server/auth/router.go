package auth

import (
	"github.com/go-chi/chi/v5"
)

func BuildRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/login", Login)
	r.Post("/register", Register)
	r.Post("/delete", DeleteAccount)

	return r
}
