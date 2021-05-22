package v1

import (
	"context"
	"net/http"
	"strings"
)

func (api *Api) auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			errorResponse("authorization failed", http.StatusUnauthorized)
			return
		}

		client, err := api.Clients.Get(auth[1])
		if err != nil {
			return
		}

		ctx := context.WithValue(r.Context(), "client", client)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}