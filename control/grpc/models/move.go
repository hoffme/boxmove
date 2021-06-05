package models

import (
	"context"

	"github.com/hoffme/boxmove/app"

	"github.com/hoffme/boxmove/controls/grpc/proto/move"
)

type MoveProtoService struct {
	move.ServiceServer
	app *app.Service
}

func NewMoveProtoService(app *app.Service) move.ServiceServer {
	return &MoveProtoService{
		app: app,
	}
}

func (m *MoveProtoService) GetAll(ctx context.Context, request *move.FilterRequest) (*move.ListMovesResponse, error) {
	panic("implement me")
}

func (m *MoveProtoService) Get(ctx context.Context, request *move.IdRequest) (*move.MoveResponse, error) {
	panic("implement me")
}

func (m *MoveProtoService) Create(ctx context.Context, request *move.CreateRequest) (*move.MoveResponse, error) {
	panic("implement me")
}
