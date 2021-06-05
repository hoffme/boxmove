package models

import (
	"context"

	"github.com/hoffme/boxmove/app"

	"github.com/hoffme/boxmove/interface/grpc/proto/box"
)

type BoxProtoService struct {
	box.ServiceServer
	app *app.Service
}

func NewBoxProtoService(app *app.Service) box.ServiceServer {
	return &BoxProtoService{
		app: app,
	}
}

func (b *BoxProtoService) Parent(ctx context.Context, request *box.IdRequest) (*box.BoxResponse, error) {
	panic("implement me")
}

func (b *BoxProtoService) Ancestors(ctx context.Context, request *box.IdRequest) (*box.ListBoxesResponse, error) {
	panic("implement me")
}

func (b *BoxProtoService) Children(ctx context.Context, request *box.IdRequest) (*box.ListBoxesResponse, error) {
	panic("implement me")
}

func (b *BoxProtoService) Decedents(ctx context.Context, request *box.IdRequest) (*box.ListBoxesResponse, error) {
	panic("implement me")
}

func (b *BoxProtoService) GetAll(ctx context.Context, request *box.FilterRequest) (*box.ListBoxesResponse, error) {
	panic("implement me")
}

func (b *BoxProtoService) Get(ctx context.Context, request *box.IdRequest) (*box.BoxResponse, error) {
	panic("implement me")
}

func (b *BoxProtoService) Create(ctx context.Context, request *box.CreateRequest) (*box.BoxResponse, error) {
	panic("implement me")
}

func (b *BoxProtoService) Update(ctx context.Context, request *box.UpdateRequest) (*box.BoxResponse, error) {
	panic("implement me")
}

func (b *BoxProtoService) Delete(ctx context.Context, request *box.IdRequest) (*box.BoxResponse, error) {
	panic("implement me")
}

func (b *BoxProtoService) Restore(ctx context.Context, request *box.IdRequest) (*box.BoxResponse, error) {
	panic("implement me")
}

func (b *BoxProtoService) Remove(ctx context.Context, request *box.IdRequest) (*box.BoxResponse, error) {
	panic("implement me")
}
