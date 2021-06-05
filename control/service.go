package control

import (
	"log"
	"sync"

	"github.com/hoffme/boxmove/app"

	"github.com/hoffme/boxmove/controls/grpc"
	"github.com/hoffme/boxmove/controls/web"
)

type server interface {
	Start() error
}

type Service struct {
	servers []server
	wg 		sync.WaitGroup
}

func NewService(app *app.Service) (*Service, error) {
	service := &Service{
		servers: []server{
			web.New(app),
			grpc.New(app),
		},
	}

	return service, nil
}

func (s *Service) Start() {
	for _, srv := range s.servers {
		go s.StartServer(srv)
	}

	s.wg.Wait()
}

func (s *Service) StartServer(srv server) {
	s.wg.Add(1)
	defer s.wg.Done()

	err := srv.Start()
	if err != nil {
		log.Fatal(err)
	}
}