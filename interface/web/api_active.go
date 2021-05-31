package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *API) routerActives(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
	r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {})
	r.Put("/{id}/delete", func(w http.ResponseWriter, r *http.Request) {})
	r.Put("/{id}/restore", func(w http.ResponseWriter, r *http.Request) {})
	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {})
}

