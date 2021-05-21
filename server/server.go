package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	addr   string
	router http.Handler
}

func New(addr string, router http.Handler) *Server {
	return &Server{
		addr: addr,
		router: router,
	}
}

func (s *Server) Run() error {
	fmt.Printf("Starting on %s\n", s.addr)
	return http.ListenAndServe(s.addr, s.router)
}
