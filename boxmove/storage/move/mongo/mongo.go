package mongo

import (
	"context"

	"github.com/hoffme/boxmove/boxmove/move"

	"github.com/hoffme/boxmove/storage"
)

type MongoStorage struct {
	conn           *storage.Connection
	ctx            context.Context
	collectionName string
	key 		   string
}

func NewMongoStorage(conn *storage.Connection, collectionName, key string) (move.Storage, error) {
	return &MongoStorage{
		conn:           conn,
		ctx:            conn.Ctx,
		collectionName: collectionName,
		key: 			key,
	}, nil
}
