package clients

import (
	"context"

	"github.com/hoffme/boxmove/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Clients struct {
	conn 	       *storage.Connection
	ctx 	       context.Context
	collectionName string
}

func New(conn *storage.Connection) *Clients {
	return &Clients{
		conn: conn,
		ctx: conn.Ctx,
		collectionName: "clients",
	}
}

func (c *Clients) collection() *mongo.Collection {
	return c.conn.DB().Collection(c.collectionName)
}

func (c *Clients) New() (*Client, error) {
	client := newClient()

	result, err := c.collection().InsertOne(c.ctx, client)
	if err != nil {
		return nil, err
	}

	client.Id = result.InsertedID.(primitive.ObjectID)

	err = client.Init(c.conn)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Clients) Get(key string) (*Client, error) {
	idP, err := primitive.ObjectIDFromHex(key)
	if err != nil {
		return nil, err
	}

	client := &Client{ Id: idP }

	err = c.collection().FindOne(c.ctx, bson.M{ "_id": idP }).Decode(client)
	if err != nil {
		return nil, err
	}

	err = client.Init(c.conn)
	if err != nil {
		return nil, err
	}

	return client, nil
}