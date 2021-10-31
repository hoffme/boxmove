package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type conn struct {
	Client       *mongo.Client
	Ctx          context.Context
	DatabaseName string
}

func connect(uri, database string, ctx context.Context) (*conn, error) {
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

	store := &conn{
		Client:       client,
		Ctx:          ctx,
		DatabaseName: database,
	}

	return store, nil
}

func (c *conn) DB() *mongo.Database {
	return c.Client.Database(c.DatabaseName)
}

func (c *conn) Collection(name string) *mongo.Collection {
	return c.Client.Database(c.DatabaseName).Collection(name)
}

func (c *conn) Close() error {
	return c.Client.Disconnect(c.Ctx)
}
