package server

import (
	"log"

	"github.com/hoffme/boxmove/app"
	"github.com/hoffme/boxmove/server/web"
	"github.com/hoffme/boxmove/utils"
)

type Service struct {
	http *web.ServerHTTP
}

func NewService(app *app.Service) (*Service, error) {
	httpAddr   := utils.GetEnv("HTTP_ADDR", ":3000")
	httpRouter := web.Router(app)
	httpServer := web.NewServerHTTP(httpAddr, httpRouter)

	service := &Service{
		http: httpServer,
	}

	return service, nil
}

func (s *Service) Start() {
	err := s.http.Run()
	if err != nil {
		log.Fatal(err)
	}
}