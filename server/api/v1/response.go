package v1

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Success bool        `json:"success"`
	Status  int         `json:"status,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func successResponse(result interface{}, status int) *response {
	return &response{
		Success: true,
		Status:  status,
		Result:  result,
	}
}

func errorResponse(error interface{}, status int) *response {
	return &response{
		Success: true,
		Status:  status,
		Error:   error,
	}
}

func (r *response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	_ = json.NewEncoder(w).Encode(r)
}
