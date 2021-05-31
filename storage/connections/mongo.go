package connections

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	Client   	 *mongo.Client
	Ctx      	 context.Context
	DatabaseName string
}

func MongoDB(uri, database string, ctx context.Context) (*MongoConnection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	store := &MongoConnection{
		Client:       client,
		Ctx:          ctx,
		DatabaseName: database,
	}

	return store, nil
}

func (c *MongoConnection) DB() *mongo.Database {
	return c.Client.Database(c.DatabaseName)
}

func (c *MongoConnection) Close() error {
	return c.Client.Disconnect(c.Ctx)
}