package app

import (
	"github.com/hoffme/boxmove/pkg/app/box"
	"github.com/hoffme/boxmove/pkg/app/item"
	"github.com/hoffme/boxmove/pkg/app/move"

	"github.com/hoffme/boxmove/pkg/storage"
)

type App interface {
	SetStorage(storage storage.Storage) error

	FindBox(id string) (box.Box, error)
	SearchBox(filter *box.Filter) ([]box.Box, error)
	CreateBox(params *box.CreateParams) (box.Box, error)

	FindItem(id string) (item.Item, error)
	SearchItem(filter *item.Filter) ([]item.Item, error)
	CreateItem(params *item.CreateParams) (item.Item, error)

	FindMove(id string) (move.Move, error)
	SearchMove(filter *move.Filter) ([]move.Move, error)
	CreateMove(params *move.CreateParams) (move.Move, error)
}
