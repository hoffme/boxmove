package client

import (
	"time"
)

func New(store Storage, params *CreateClientParams) (*Client, error) {
	dto, err := store.New(&DTOCreateParams{
		Name: params.Name,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		dto: dto,
		store: store,
	}, nil
}

func Get(store Storage, id string) (*Client, error) {
	dto, err := store.Get(id)
	if err != nil {
		return nil, err
	}

	return &Client{
		dto: dto,
		store: store,
	}, err
}