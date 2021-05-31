package web

import (
	"context"
	"net/http"
	"time"

	"github.com/hoffme/boxmove/app"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type API struct {
	App *app.Service
}

func NewApi(app *app.Service) *API {
	return &API{ App: app }
}

func (a *API) Router() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(10 * time.Second))

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	})

	router.Route("/clients", a.routerClients)

	router.Route("/boxmove/{client}", func(r chi.Router) {
		r.Use(a.clientMiddleware)

		r.Route("/actives", a.routerActives)
		r.Route("/boxes", a.routerBoxes)
		r.Route("/moves", a.routerMoves)
	})

	return router
}

func (a *API) clientMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "client")

		client, err := a.App.Clients.FindOne(id)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if client == nil {
			http.Error(w, "client not found", 404)
			return
		}

		ctx := context.WithValue(r.Context(), "client", client)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}