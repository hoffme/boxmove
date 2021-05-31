package mongo

import (
	"context"

	"github.com/hoffme/boxmove/storage"
)

type Storage struct {
	connection     *storage.MongoConnection
	ctx            context.Context
	collectionName string
}

func New(conn *storage.MongoConnection, collectionName string) (*Storage, error) {
	storageMongo := &Storage{
		connection:     conn,
		ctx:            conn.Ctx,
		collectionName: collectionName,
	}

	return storageMongo, nil
}