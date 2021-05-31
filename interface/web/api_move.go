package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *API) routerMoves(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
}
