package api

import (
	"context"
	"net/http"

	"github.com/hoffme/boxmove/clients"

	"github.com/go-chi/chi/v5"
)

func authMiddleware(cls *clients.Clients) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := chi.URLParam(r, "key")

			client, err := cls.Get(key)
			if err != nil {
				return
			}

			ctx := context.WithValue(r.Context(), "client", client)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
