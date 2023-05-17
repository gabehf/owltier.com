package list

import "github.com/go-chi/chi/v5"

func BuildRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/new", NewList)
	r.Post("/delete", DeleteList) // should this even exist?
	r.Get("/{id}", GetList)

	return r
}
