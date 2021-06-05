package _interface

import (
	"log"

	"github.com/hoffme/boxmove/app"

	"github.com/hoffme/boxmove/interface/grpc"
)

type server interface {
	Start() error
}

type Service struct {
	servers []server
}

func NewService(app *app.Service) (*Service, error) {
	service := &Service{
		servers: []server{
			grpc.New(app),
		},
	}

	return service, nil
}

func (s *Service) Start() {
	// wg := sync.WaitGroup{}

	// wg.Add(len(s.servers))

	// for _, srv := range s.servers {
	// 	go func(srv server) {
	// 		defer wg.Done()

	// 		err := srv.Start()
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	}(srv)
	// }

	// wg.Wait()

	err := s.servers[0].Start()
	if err != nil {
		log.Fatal(err)
	}
}
