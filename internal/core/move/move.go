package move

import (
	"github.com/hoffme/boxmove/pkg/core/move"

	"github.com/hoffme/boxmove/pkg/storage"
	storageItem "github.com/hoffme/boxmove/pkg/storage/move"
)

type move_e struct {
	storage storage.Storage
	id      string
	dto     *storageItem.DTO
}

func newMove(storage storage.Storage, dto *storageItem.DTO) move.Move {
	return &move_e{
		storage: storage,
		id:      dto.ID,
		dto:     dto,
	}
}

func (b *move_e) refresh() error {
	dto, err := b.storage.Move().Find(b.id)
	if err != nil {
		return err
	}

	b.dto = dto

	return nil
}

func (b *move_e) ID() string {
	return b.id
}

func (b *move_e) Data() (*move.Data, error) {
	if b.dto == nil {
		err := b.refresh()
		if err != nil {
			return nil, err
		}
	}

	data := &move.Data{
		FromID:   b.dto.FromID,
		ToID:     b.dto.ToID,
		Date:     b.dto.Date,
		Quantity: b.dto.Quantity,
		ItemID:   b.dto.ItemID,
		Meta:     b.dto.Meta,
	}

	return data, nil
}
