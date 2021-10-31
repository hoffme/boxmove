package box

import (
	"github.com/hoffme/boxmove/pkg/core/box"

	"github.com/hoffme/boxmove/pkg/storage"
	boxStorage "github.com/hoffme/boxmove/pkg/storage/box"
)

type store struct {
	storage storage.Storage
}

func NewStore(storage storage.Storage) box.Store {
	return &store{storage: storage}
}

func (m *store) Find(id string) (box.Box, error) {
	dto, err := m.storage.Box().Find(id)
	if err != nil {
		return nil, err
	}

	return newBox(m.storage, dto), nil
}

func (m *store) Search(params *box.StoreSearchParams) ([]box.Box, error) {
	dtos, err := m.storage.Box().Search(&boxStorage.Filter{
		ID:         params.ID,
		Name:       params.Name,
		ParentID:   params.ParentID,
		AncestorID: params.AncestorID,
	})
	if err != nil {
		return nil, err
	}

	boxes := []box.Box{}

	for _, dto := range dtos {
		boxes = append(boxes, newBox(m.storage, dto))
	}

	return boxes, nil
}

func (m *store) Create(params *box.StoreCreateParams) (box.Box, error) {
	route := []string{}

	if len(params.Parent) > 0 {
		_, err := m.Find(params.Parent)
		if err != nil {
			return nil, err
		}

		route = append(route, params.Parent)
	}

	dto, err := m.storage.Box().Create(&boxStorage.Fields{
		Name:  params.Name,
		Route: route,
		Meta:  params.Meta,
	})
	if err != nil {
		return nil, err
	}

	return newBox(m.storage, dto), nil
}
