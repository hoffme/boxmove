package item

import (
	"github.com/hoffme/boxmove/pkg/core/item"

	"github.com/hoffme/boxmove/pkg/storage"
	itemStorage "github.com/hoffme/boxmove/pkg/storage/item"
)

type store struct {
	storage storage.Storage
}

func NewStore(storage storage.Storage) item.Store {
	return &store{storage: storage}
}

func (m *store) Find(id string) (item.Item, error) {
	dto, err := m.storage.Item().Find(id)
	if err != nil {
		return nil, err
	}

	return newItem(m.storage, dto), nil
}

func (m *store) Search(params *item.StoreSearchParams) ([]item.Item, error) {
	dtos, err := m.storage.Item().Search(&itemStorage.Filter{
		ID:   params.ID,
		Name: params.Name,
	})
	if err != nil {
		return nil, err
	}

	items := []item.Item{}

	for _, dto := range dtos {
		items = append(items, newItem(m.storage, dto))
	}

	return items, nil
}

func (m *store) Create(params *item.StoreCreateParams) (item.Item, error) {
	dto, err := m.storage.Item().Create(&itemStorage.Fields{
		Name: params.Name,
		Meta: params.Meta,
	})
	if err != nil {
		return nil, err
	}

	return newItem(m.storage, dto), nil
}
