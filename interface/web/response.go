package web

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Result interface{} `json:"result,omitempty"`
	Error  interface{} `json:"error,omitempty"`
	Status int 		   `json:"status"`
}

func (r response) Send(w http.ResponseWriter) {
	if r.Status == 0 {
		r.Status = 200
	}

	w.WriteHeader(r.Status)
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(r)
}
