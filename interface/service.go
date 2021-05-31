package _interface

import (
	"log"

	"github.com/hoffme/boxmove/app"
	"github.com/hoffme/boxmove/interface/web"
	"github.com/hoffme/boxmove/utils"
)

type Service struct {
	http *web.ServerHTTP
}

func NewService(app *app.Service) (*Service, error) {
	httpAddr := utils.GetEnv("HTTP_ADDR", ":3000")
	api 	 := web.NewApi(app)

	httpServer := web.NewServerHTTP(httpAddr, api.Router())

	service := &Service{
		http: httpServer,
	}

	return service, nil
}

func (s *Service) Start() {
	err := s.http.Start()
	if err != nil {
		log.Fatal(err)
	}
}