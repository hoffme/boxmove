package api

import (
	"net/http"

	"github.com/hoffme/boxmove/clients"

	"github.com/go-chi/chi/v5"
)

func CreateRouter(cls *clients.Clients) http.Handler {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		successResponse("pong", http.StatusOK).Send(w)
	})

	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		client, err := cls.New()

		if err != nil {
			errorResponse(err.Error(), http.StatusNotFound).Send(w)
			return
		}

		successResponse(client, http.StatusOK).Send(w)
	})

	r.Route("/key/{key}", func(r chi.Router) {
		r.Use(authMiddleware(cls))


	})

	return r
}