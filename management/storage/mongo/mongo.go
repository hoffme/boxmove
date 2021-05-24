package mongo

import (
	"context"

	"github.com/hoffme/boxmove/management/client"

	"github.com/hoffme/boxmove/storage"
)

type Storage struct {
	conn           *storage.Connection
	ctx            context.Context
	collectionName string
}

func NewMongoStorage(conn *storage.Connection, collectionName string) client.Storage {
	return &Storage{
		conn: 			conn,
		ctx:            conn.Ctx,
		collectionName: collectionName,
	}
}

