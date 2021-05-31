package web

import (
	"encoding/json"
	"net/http"

	"github.com/hoffme/boxmove/boxmove/client"

	"github.com/go-chi/chi/v5"
)

func (a *API) routerClients(r chi.Router) {
	r.Post("/search", a.clientSearch)
	r.Get("/", a.clientGetAll)
	r.Get("/{id}", a.clientGetById)
	r.Post("/", a.clientCreate)
	r.Put("/{id}", a.clientUpdate)
	r.Put("/{id}/delete", a.clientDelete)
	r.Put("/{id}/restore", a.clientRestore)
	r.Delete("/{id}", a.clientRemove)
}

func (a *API) clientSearch(w http.ResponseWriter, r *http.Request) {
	var filter client.Filter

	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		response{ Error: err.Error(), Status: http.StatusBadRequest }.Send(w)
		return
	}

	clients, err := a.App.Clients.FindAll(&filter)
	if err != nil {
		response{ Error: err.Error(), Status: http.StatusBadRequest }.Send(w)
		return
	}

	result := make([]client.View, len(clients))
	for _, clt := range clients {
		result = append(result, *clt.View())
	}

	response{ Result: result }.Send(w)
}

func (a *API) clientGetAll(w http.ResponseWriter, r *http.Request) {
	clients, err := a.App.Clients.FindAll(&client.Filter{})
	if err != nil {
		response{ Error: err.Error(), Status: http.StatusBadRequest }.Send(w)
		return
	}

	result := make([]client.View, len(clients))
	for _, clt := range clients {
		result = append(result, *clt.View())
	}

	response{ Result: result }.Send(w)
}

func (a *API) clientGetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	result, err := a.App.Clients.FindOne(id)
	if err != nil {
		response{ Error: err.Error(), Status: http.StatusBadRequest }.Send(w)
		return
	}

	response{ Result: result.View() }.Send(w)
}

func (a *API) clientCreate(w http.ResponseWriter, r *http.Request) {
	var params client.CreateParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		response{ Error: err.Error(), Status: http.StatusBadRequest }.Send(w)
		return
	}

	result, err := a.App.Clients.New(&params)
	if err != nil {
		response{ Error: err.Error(), Status: http.StatusBadRequest }.Send(w)
		return
	}

	response{ Result: result.View() }.Send(w)
}

func (a *API) clientUpdate(w http.ResponseWriter, r *http.Request) {
	var params client.UpdateParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		response{ Error: err.Error(), Status: http.StatusBadRequest }.Send(w)
		return
	}

	id := chi.URLParam(r, "id")

	clt, err := a.App.Clients.FindOne(id)
	if err != nil {
		response{ Error: err.Error(), Status: http.StatusBadRequest }.Send(w)
		return
	}

	err = clt.Update(&params)
	if err != nil {
		response{ Error: err.Error(), Status: http.StatusBadRequest }.Send(w)
		return
	}


	response{ Result: clt.View() }.Send(w)
}

func (a *API) clientDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	clt, err := a.App.Clients.FindOne(id)
	if err != nil {
		response{ Error: err.Error(), Status: http.StatusBadRequest }.Send(w)
		return
	}

	err = clt.Delete()
	if err != nil {
		response{ Error: err.Error(), Status: http.StatusBadRequest }.Send(w)
		return
	}


	response{ Result: clt.View() }.Send(w)
}

func (a *API) clientRestore(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	clt, err := a.App.Clients.FindOne(id)
	if err != nil {
		response{ Error: err.Error(), Status: http.StatusBadRequest }.Send(w)
		return
	}

	err = clt.Restore()
	if err != nil {
		response{ Error: err.Error(), Status: http.StatusBadRequest }.Send(w)
		return
	}


	response{ Result: clt.View() }.Send(w)
}

func (a *API) clientRemove(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	clt, err := a.App.Clients.FindOne(id)
	if err != nil {
		response{ Error: err.Error(), Status: http.StatusBadRequest }.Send(w)
		return
	}

	err = clt.Remove()
	if err != nil {
		response{ Error: err.Error(), Status: http.StatusBadRequest }.Send(w)
		return
	}


	response{ Result: clt.View() }.Send(w)
}