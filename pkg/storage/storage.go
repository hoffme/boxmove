package storage

import (
	"github.com/hoffme/boxmove/pkg/storage/box"
	"github.com/hoffme/boxmove/pkg/storage/item"
	"github.com/hoffme/boxmove/pkg/storage/move"
)

type Stores struct {
	Items item.Store
	Boxes box.Store
	Moves move.Store
}

type Storage interface {
	Stores() *Stores
	Close() error
}
