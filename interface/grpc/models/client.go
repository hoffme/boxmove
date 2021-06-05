package models

import (
	"context"

	"github.com/hoffme/boxmove/app"

	"github.com/hoffme/boxmove/interface/grpc/proto/client"
)

type ClientProtoService struct {
	client.ServiceServer
	app *app.Service
}

func NewClientProtoService(app *app.Service) client.ServiceServer {
	return &ClientProtoService{
		app: app,
	}
}

func (s *ClientProtoService) GetAll(ctx context.Context, params *client.FilterRequest) (*client.ListClientsResponse, error) {
	panic("implement me")
}

func (s *ClientProtoService) Get(ctx context.Context, id *client.IdRequest) (*client.ClientResponse, error) {
	panic("implement me")
}

func (s *ClientProtoService) Create(ctx context.Context, params *client.CreateRequest) (*client.ClientResponse, error) {
	panic("implement me")
}

func (s *ClientProtoService) Update(ctx context.Context, params *client.UpdateRequest) (*client.ClientResponse, error) {
	panic("implement me")
}

func (s *ClientProtoService) Delete(ctx context.Context, id *client.IdRequest) (*client.ClientResponse, error) {
	panic("implement me")
}

func (s *ClientProtoService) Restore(ctx context.Context, id *client.IdRequest) (*client.ClientResponse, error) {
	panic("implement me")
}

func (s *ClientProtoService) Remove(ctx context.Context, id *client.IdRequest) (*client.ClientResponse, error) {
	panic("implement me")
}
