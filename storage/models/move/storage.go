package move

import (
	"context"
	"github.com/hoffme/boxmove/storage/connections"
)

type Storage struct {
	connection     *connections.MongoConnection
	ctx            context.Context
	collectionName string
}

func New(conn *connections.MongoConnection, collectionName string) (*Storage, error) {
	storageMongo := &Storage{
		connection:     conn,
		ctx:            conn.Ctx,
		collectionName: collectionName,
	}

	return storageMongo, nil
}