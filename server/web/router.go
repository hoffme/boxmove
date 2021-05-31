package web

import (
	"net/http"
	"time"

	"github.com/hoffme/boxmove/app"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router(app *app.Service) http.Handler {
	api := api{ App: app }

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

	router.Route("/clients", api.routerClients)

	router.Route("/boxmove/{client}", func(r chi.Router) {
		r.Use(api.middlewareClient)

		r.Route("/actives", api.routerActives)
		r.Route("/boxes", api.routerBoxes)
		r.Route("/moves", api.routerMoves)
	})

	return router
}