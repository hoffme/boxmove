package management

import (
	"github.com/hoffme/boxmove/management/client"
	clientStore "github.com/hoffme/boxmove/management/storage/mongo"
	"github.com/hoffme/boxmove/storage"
)

type Management struct {
	clients client.Storage
}

func New(conn *storage.Connection) *Management {
	return &Management{
		clients: clientStore.NewMongoStorage(conn, "clients"),
	}
}

func (m *Management) NewClient(params *client.CreateClientParams) (*client.Client, error) {
	return client.New(m.clients, params)
}

func (m *Management) GetClient(id string) (*client.Client, error) {
	return client.Get(m.clients, id)
}

func (m *Management) RemoveClient(id string) error {
	cli, err := client.Get(m.clients, id)
	if err != nil {
		return err
	}

	return cli.Delete()
}
