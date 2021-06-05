package grpc

import (
	"log"
	"net"

	"github.com/hoffme/boxmove/app"
	"github.com/hoffme/boxmove/utils"

	"github.com/hoffme/boxmove/interface/grpc/models"

	"github.com/hoffme/boxmove/interface/grpc/proto/box"
	"github.com/hoffme/boxmove/interface/grpc/proto/client"
	"github.com/hoffme/boxmove/interface/grpc/proto/move"

	"google.golang.org/grpc"
)

type Server struct {
	addr    string
	network string
	server  *grpc.Server
}

func New(app *app.Service) *Server {
	addr := utils.GetEnv("GRPC_ADDR", ":5000")
	network := utils.GetEnv("GRPC_NET", "tcp")

	server := &Server{
		addr:    addr,
		network: network,
	}

	server.loadServices(app)

	return server
}

func (s *Server) Start() error {
	log.Printf("Starting grpc %s on %s\n", s.network, s.addr)

	lis, err := net.Listen(s.network, s.addr)
	if err != nil {
		return err
	}

	return s.server.Serve(lis)
}

func (s *Server) loadServices(app *app.Service) {
	server := grpc.NewServer()

	clientService := models.NewClientProtoService(app)
	client.RegisterServiceServer(server, clientService)

	boxService := models.NewBoxProtoService(app)
	box.RegisterServiceServer(server, boxService)

	moveService := models.NewMoveProtoService(app)
	move.RegisterServiceServer(server, moveService)

	s.server = server
}
