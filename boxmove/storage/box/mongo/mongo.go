package mongo

import (
	"context"

	"github.com/hoffme/boxmove/boxmove/box"
	"github.com/hoffme/boxmove/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStorage struct {
	conn           *storage.Connection
	ctx            context.Context
	collectionName string
	key 		   string
}

func NewMongoStorage(conn *storage.Connection, collectionName, key string) (box.Storage, error) {
	store := &MongoStorage{
		conn: 			conn,
		ctx:            conn.Ctx,
		collectionName: collectionName,
		key: 			key,
	}

	_, err := store.collection().Indexes().CreateOne(store.ctx, mongo.IndexModel{
		Keys: bson.D{ { "name", "text" } }, Options: nil,
	})
	if err != nil {
		return nil, err
	}

	return store, nil
}

