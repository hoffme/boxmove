package move

import (
	"time"

	"github.com/hoffme/boxmove/pkg/core/move"

	"github.com/hoffme/boxmove/pkg/storage"
	moveStorage "github.com/hoffme/boxmove/pkg/storage/move"
)

type store struct {
	storage storage.Storage
}

func NewStore(storage storage.Storage) move.Store {
	return &store{storage: storage}
}

func (m *store) Find(id string) (move.Move, error) {
	dto, err := m.storage.Move().Find(id)
	if err != nil {
		return nil, err
	}

	return newMove(m.storage, dto), nil
}

func (m *store) Search(params *move.StoreSearchParams) ([]move.Move, error) {
	dtos, err := m.storage.Move().Search(&moveStorage.Filter{
		ID:          params.ID,
		ItemID:      params.ItemID,
		FromID:      params.FromID,
		ToID:        params.ToID,
		DateFrom:    params.DateFrom,
		DateTo:      params.DateTo,
		QuantityMin: params.QuantityMin,
		QuantityMax: params.QuantityMax,
	})
	if err != nil {
		return nil, err
	}

	moves := []move.Move{}

	for _, dto := range dtos {
		moves = append(moves, newMove(m.storage, dto))
	}

	return moves, nil
}

func (m *store) Create(params *move.StoreCreateParams) (move.Move, error) {
	dto, err := m.storage.Move().Create(&moveStorage.Fields{
		Date:     time.Now(),
		FromID:   params.FromID,
		ToID:     params.ToID,
		ItemID:   params.ItemID,
		Quantity: params.Quantity,
		Meta:     params.Meta,
	})
	if err != nil {
		return nil, err
	}

	return newMove(m.storage, dto), nil
}
