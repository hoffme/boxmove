package clients

import (
	"time"

	"github.com/hoffme/boxmove/boxmove"
	"github.com/hoffme/boxmove/storage"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	Id           primitive.ObjectID  `json:"-" bson:"_id,omitempty"`
	CreatedAt    time.Time           `json:"created_at" bson:"created_at"`
	Key          string              `json:"key" bson:"-"`
	Manager      *boxmove.Management `json:"-" bson:"-"`
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