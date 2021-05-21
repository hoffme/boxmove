package clients

import (
	"time"

	"github.com/hoffme/boxmove/boxmove"
	"github.com/hoffme/boxmove/storage"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	Id           primitive.ObjectID  `bson:"_id,omitempty"`
	CreatedAt    time.Time           `bson:"created_at"`
	Key          string              `bson:"-"`
	Manager      *boxmove.Management `bson:"-"`
}

func newClient() *Client {
	return &Client{ CreatedAt: time.Now() }
}

func (c *Client) Init(conn *storage.Connection) error {
	c.Key = c.Id.Hex()

	manager, err := boxmove.New(conn, c.Key)
	if err != nil {
		return err
	}

	c.Manager = manager

	return nil
}