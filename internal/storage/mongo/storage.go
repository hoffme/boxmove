package mongo

import (
	"context"

	"github.com/hoffme/boxmove/pkg/storage"

	"github.com/hoffme/boxmove/internal/storage/connections"
)

type Storage struct {
	mongoConnection *connections.MongoConnection
	stores          *storage.Stores
}

func New(uri, database string) (storage.Storage, error) {
	storage := &Storage{}

	err := storage.setConnections(uri, database)
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *Storage) setConnections(uri, database string) error {
	ctx := context.Background()

	mongo, err := connections.MongoDB(uri, database, ctx)
	if err != nil {
		return err
	}

	s.mongoConnection = mongo

	return nil
}

func (s *Storage) Stores() *storage.Stores {
	return s.stores
}

func (s *Storage) Close() error {
	return s.mongoConnection.Close()
}
