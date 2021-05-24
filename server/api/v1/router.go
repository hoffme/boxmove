package v1

import (
	"encoding/json"
	"errors"
	"github.com/hoffme/boxmove/management/client"
	"net/http"

	"github.com/hoffme/boxmove/management"

	"github.com/go-chi/chi/v5"
)

type Api struct {
	Management *management.Management
}

func (api *Api) Router(r chi.Router)  {
	r.Post("/client", api.newClientHandler)
	r.Get("/client/:id", api.newClientHandler)
	r.Delete("/client/:id", api.deleteClientHandler)

	r.Route("/boxmove", api.boxmoveRouter)
}

func (api *Api) newClientHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer func() { _ = body.Close() }()

	params := client.CreateClientParams{}
	err := json.NewDecoder(body).Decode(&params)
	if err != nil {
		errorResponse(err.Error(), http.StatusBadRequest).Send(w)
		return
	}

	clt, err := api.Management.NewClient(&params)
	if err != nil {
		errorResponse(err.Error(), http.StatusBadRequest).Send(w)
		return
	}

	successResponse(clt.View(), http.StatusOK).Send(w)
}

func (api *Api) getClientHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(id) == 0 {
		err := errors.New("invalid id")
		errorResponse(err.Error(), http.StatusBadRequest).Send(w)
		return
	}

	clt, err := api.Management.GetClient(id)
	if err != nil {
		errorResponse(err.Error(), http.StatusBadRequest).Send(w)
		return
	}

	successResponse(clt.View(), http.StatusOK).Send(w)
}

func (api *Api) deleteClientHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(id) == 0 {
		err := errors.New("invalid id")
		errorResponse(err.Error(), http.StatusBadRequest).Send(w)
		return
	}

	clt, err := api.Management.GetClient(id)
	if err != nil {
		errorResponse(err.Error(), http.StatusBadRequest).Send(w)
		return
	}

	err = clt.Delete()
	if err != nil {
		errorResponse(err.Error(), http.StatusBadRequest).Send(w)
		return
	}

	successResponse(clt.View(), http.StatusOK).Send(w)
}

func (api *Api) boxmoveRouter(r chi.Router) {
	r.Use(api.auth)
}