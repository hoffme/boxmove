package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *API) routerBoxes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/{id}/parent", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/{id}/ancestors", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/{id}/children", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/{id}/decedents", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/{id}/decedents/tree", func(w http.ResponseWriter, r *http.Request) {})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
	r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {})
	r.Put("/{id}/delete", func(w http.ResponseWriter, r *http.Request) {})
	r.Put("/{id}/restore", func(w http.ResponseWriter, r *http.Request) {})
	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {})
}