package models

import (
	"context"

	"github.com/hoffme/boxmove/app"
	clientApp "github.com/hoffme/boxmove/boxmove/client"

	"github.com/hoffme/boxmove/controls/grpc/proto/client"
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
	filter := &clientApp.Filter{
		ID: params.Ids,
		Name: params.Name,
		Deleted: params.Deleted,
	}

	clients, err := s.app.Clients.FindAll(filter)
	if err != nil {
		return nil, err
	}

	result := []client.Client{}

	for _, clientModel := range clients {
		view := clientModel.View()
		if view == nil {
			continue
		}

		result = append(result, client.Client{
			Id: view.ID,
			Name: view.Name,
			CreatedAt: view.CreatedAt,
			UpdatedAt: view.UpdatedAt,
			DeletedAt: view.DeletedAt,
		})
	}

	return &client.ListClientsResponse{ Clients: result }, nil
}

func (s *ClientProtoService) Get(ctx context.Context, id *client.IdRequest) (*client.ClientResponse, error) {
	return &client.ClientResponse{}, nil
}

func (s *ClientProtoService) Create(ctx context.Context, params *client.CreateRequest) (*client.ClientResponse, error) {
	return &client.ClientResponse{}, nil
}

func (s *ClientProtoService) Update(ctx context.Context, params *client.UpdateRequest) (*client.ClientResponse, error) {
	return &client.ClientResponse{}, nil
}

func (s *ClientProtoService) Delete(ctx context.Context, id *client.IdRequest) (*client.ClientResponse, error) {
	return &client.ClientResponse{}, nil
}

func (s *ClientProtoService) Restore(ctx context.Context, id *client.IdRequest) (*client.ClientResponse, error) {
	return &client.ClientResponse{}, nil
}

func (s *ClientProtoService) Remove(ctx context.Context, id *client.IdRequest) (*client.ClientResponse, error) {
	return &client.ClientResponse{}, nil
}