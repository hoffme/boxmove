package storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Client *mongo.Client
	Ctx    context.Context
}

var connection *Connection

func GetConnection(uri string, ctx context.Context) (*Connection, error) {
	if connection != nil {
		return connection, nil
	}

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

	connection = &Connection{
		Client: client,
		Ctx: ctx,
	}

	return connection, nil
}

func (c *Connection) Close() {
	err := c.Client.Disconnect(c.Ctx)
	if err != nil {
		log.Fatal(err)
	}
}