package mongo

import (
	"context"

	"github.com/hoffme/boxmove/pkg/storage"
	"github.com/hoffme/boxmove/pkg/storage/box"
	"github.com/hoffme/boxmove/pkg/storage/item"
	"github.com/hoffme/boxmove/pkg/storage/move"
)

type Storage struct {
	connection *conn

	box  *box_s
	move *move_s
	item *item_s
}

func New(settings *Settings) (storage.Storage, error) {
	var err error

	storage := &Storage{}

	ctx := context.Background()

	storage.connection, err = connect(settings.Uri, settings.Database, ctx)
	if err != nil {
		return nil, err
	}

	storage.box = newBoxStore(storage.connection, settings.BoxCollectionName)
	storage.move = newMoveStore(storage.connection, settings.MoveCollectionName)
	storage.item = newItemStore(storage.connection, settings.ItemCollectionName)

	return storage, nil
}

func (s *Storage) Box() box.Storage {
	return s.box
}

func (s *Storage) Item() item.Storage {
	return s.item
}

func (s *Storage) Move() move.Storage {
	return s.move
}

func (s *Storage) Close() error {
	return s.connection.Close()
}
