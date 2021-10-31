package storage

import (
	"github.com/hoffme/boxmove/pkg/storage/box"
	"github.com/hoffme/boxmove/pkg/storage/item"
	"github.com/hoffme/boxmove/pkg/storage/move"
)

type Storage interface {
	Item() item.Storage
	Box() box.Storage
	Move() move.Storage

	Close() error
}
