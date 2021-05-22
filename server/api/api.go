package api

import (
	"net/http"
	"time"

	"github.com/hoffme/boxmove/clients"

	v1 "github.com/hoffme/boxmove/server/api/v1"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateRouter(cls *clients.Clients) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	})

	routerV1 := &v1.Api{ Clients: cls }
	r.Route("/v1", routerV1.Router)

	return r
}