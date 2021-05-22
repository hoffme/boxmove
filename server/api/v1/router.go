package v1

import (
	"github.com/hoffme/boxmove/clients"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Api struct {
	Clients *clients.Clients
}

func (api *Api) Router(r chi.Router)  {
	r.Post("/register", api.registerHandler)

	r.Route("/boxmove", func(r chi.Router) {
		r.Use(api.auth)
	})
}

func (api *Api) registerHandler(w http.ResponseWriter, r *http.Request) {
	client, err := api.Clients.New()
	if err != nil {
		errorResponse(err.Error(), http.StatusBadRequest).Send(w)
		return
	}

	successResponse(client, http.StatusOK).Send(w)
}