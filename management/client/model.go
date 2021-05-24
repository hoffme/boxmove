package client

import (
	"errors"
	"time"

	"github.com/hoffme/boxmove/boxmove/box"
	"github.com/hoffme/boxmove/boxmove/move"
	boxMongo "github.com/hoffme/boxmove/boxmove/storage/box/mongo"
	moveMongo "github.com/hoffme/boxmove/boxmove/storage/move/mongo"
	"github.com/hoffme/boxmove/storage"
)

type Client struct {
	store 		Storage
	storageBox  box.Storage
	storageMove move.Storage
	dto         DTO
}

type View struct {
	ID        string 	 `json:"id"`
	Name      string 	 `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (c *Client) Load(connection *storage.Connection) error {
	boxRepo, err := boxMongo.NewMongoStorage(connection, "boxes", c.View().ID)
	if err != nil {
		return errors.New("not created box repository: " + err.Error())
	}

	movRepo, err := moveMongo.NewMongoStorage(connection, "moves", c.View().ID)
	if err != nil {
		return errors.New("not created move repository: " + err.Error())
	}

	c.storageBox = boxRepo
	c.storageMove = movRepo

	return nil
}

func (c *Client) View() *View {
	return c.dto.View()
}

func (c *Client) Delete() error {
	return c.store.Delete(c.dto)
}