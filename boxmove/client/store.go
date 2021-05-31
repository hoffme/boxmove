package client

import (
	"errors"
)

type Store struct {
	Storage Storage
}

func (store *Store) New(params *CreateParams) (*Client, error) {
	if params == nil {
		return nil, errors.New("invalid params")
	}

	dto, err := store.Storage.Create(params)
	if err != nil {
		return nil, err
	}

	client := &Client{
		storage: store.Storage,
		id:      dto.ID(),
		dto:     dto,
	}

	return client, nil
}

func (store *Store) FindOne(id string) (*Client, error) {
	dto, err := store.Storage.FindOne(id)
	if err != nil {
		return nil, err
	}
	if dto == nil {
		return nil, errors.New("not found")
	}

	client := &Client{
		storage: store.Storage,
		id:      id,
		dto:     dto,
	}

	return client, nil
}

func (store *Store) FindAll(filter *Filter) ([]*Client, error) {
	if filter == nil {
		filter = &Filter{}
	}

	results, err := store.Storage.FindAll(filter)
	if err != nil {
		return nil, err
	}

	var objs []*Client

	for _, dto := range results {
		obj := &Client{
			storage: store.Storage,
			id:      dto.ID(),
			dto:     dto,
		}

		objs = append(objs, obj)
	}

	return objs, nil
}
