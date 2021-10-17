package box

import (
	"errors"
)

type Store struct {
	Storage Storage
}

func (store *Store) New(client string, params *CreateParams) (*Box, error) {
	if params == nil {
		return nil, errors.New("invalid params")
	}

	dto, err := store.Storage.Create(client, params)
	if err != nil {
		return nil, err
	}

	active := &Box{
		storage: store.Storage,
		id:      dto.ID(),
		dto:     dto,
		client:  client,
	}

	return active, nil
}

func (store *Store) FindOne(client, id string) (*Box, error) {
	dto, err := store.Storage.FindOne(client, id)
	if err != nil {
		return nil, err
	}
	if dto == nil {
		return nil, errors.New("not found active")
	}

	active := &Box{
		storage: store.Storage,
		id:      id,
		dto:     dto,
		client:  client,
	}

	return active, nil
}

func (store *Store) FindAll(client string, filter *Filter) ([]*Box, error) {
	if filter == nil {
		filter = &Filter{}
	}

	results, err := store.Storage.FindAll(client, filter)
	if err != nil {
		return nil, err
	}

	var objs []*Box

	for _, dto := range results {
		obj := &Box{
			storage: store.Storage,
			id:      dto.ID(),
			dto:     dto,
			client:  client,
		}

		objs = append(objs, obj)
	}

	return objs, nil
}
