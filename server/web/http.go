package web

import (
	"fmt"
	"net/http"
)

type ServerHTTP struct {
	Addr   string
	Router http.Handler
}

func NewServerHTTP(addr string, router http.Handler) *ServerHTTP {
	return &ServerHTTP{
		Addr: addr,
		Router: router,
	}
}

func (s *ServerHTTP) Run() error {
	fmt.Printf("Starting on %s\n", s.Addr)
	return http.ListenAndServe(s.Addr, s.Router)
}
