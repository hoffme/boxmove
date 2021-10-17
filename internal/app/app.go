package app

import (
	"github.com/hoffme/boxmove/pkg/app"
	"github.com/hoffme/boxmove/pkg/app/box"
	"github.com/hoffme/boxmove/pkg/app/item"
	"github.com/hoffme/boxmove/pkg/app/move"

	"github.com/hoffme/boxmove/pkg/storage"
)

type App struct {
	storage storage.Storage
}

func New() (app.App, error) {
	return &App{}, nil
}

func (a *App) SetStorage(storage storage.Storage) error {
	a.storage = storage

	return nil
}

func (a *App) FindBox(id string) (box.Box, error)
func (a *App) SearchBox(filter *box.Filter) ([]box.Box, error)
func (a *App) CreateBox(params *box.CreateParams) (box.Box, error)

func (a *App) FindItem(id string) (item.Item, error)
func (a *App) SearchItem(filter *item.Filter) ([]item.Item, error)
func (a *App) CreateItem(params *item.CreateParams) (item.Item, error)

func (a *App) FindMove(id string) (move.Move, error)
func (a *App) SearchMove(filter *move.Filter) ([]move.Move, error)
func (a *App) CreateMove(params *move.CreateParams) (move.Move, error)
