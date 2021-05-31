package web

import (
	"context"
	"net/http"

	"github.com/hoffme/boxmove/app"

	"github.com/go-chi/chi/v5"
)

type api struct {
	App *app.Service
}

func (a *api) routerClients(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
	r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {})
	r.Put("/{id}/delete", func(w http.ResponseWriter, r *http.Request) {})
	r.Put("/{id}/restore", func(w http.ResponseWriter, r *http.Request) {})
	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {})
}

func (a *api) middlewareClient(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "client")

		client, err := a.App.Clients.FindOne(id)
		if err != nil {
			http.Error(w, "client not found", 404)
			return
		}

		ctx := context.WithValue(r.Context(), "client", client)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *api) routerActives(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
	r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {})
	r.Put("/{id}/delete", func(w http.ResponseWriter, r *http.Request) {})
	r.Put("/{id}/restore", func(w http.ResponseWriter, r *http.Request) {})
	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {})
}

func (a *api) routerBoxes(r chi.Router) {
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

func (a *api) routerMoves(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
}