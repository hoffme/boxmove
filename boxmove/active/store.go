package active

import (
	"errors"
)

type Store struct {
	Storage Storage
}

func (store *Store) New(client string, params *CreateParams) (*Active, error) {
	if params == nil {
		return nil, errors.New("invalid params")
	}

	dto, err := store.Storage.Create(client, params)
	if err != nil {
		return nil, err
	}

	active := &Active{
		storage: store.Storage,
		id:      dto.ID(),
		dto:     dto,
		client:  client,
	}

	return active, nil
}

func (store *Store) FindOne(client, id string) (*Active, error) {
	dto, err := store.Storage.FindOne(client, id)
	if err != nil {
		return nil, err
	}
	if dto == nil {
		return nil, errors.New("not found active")
	}

	active := &Active{
		storage: store.Storage,
		id:      id,
		dto:     dto,
		client:  client,
	}

	return active, nil
}

func (store *Store) FindAll(client string, filter *Filter) ([]*Active, error) {
	if filter == nil {
		filter = &Filter{}
	}

	results, err := store.Storage.FindAll(client, filter)
	if err != nil {
		return nil, err
	}

	var objs []*Active

	for _, dto := range results {
		obj := &Active{
			storage: store.Storage,
			id:      dto.ID(),
			dto:     dto,
			client:  client,
		}

		objs = append(objs, obj)
	}

	return objs, nil
}
