package item

import (
	"github.com/hoffme/boxmove/pkg/core/item"

	"github.com/hoffme/boxmove/pkg/storage"
	storageItem "github.com/hoffme/boxmove/pkg/storage/item"
)

type item_e struct {
	storage storage.Storage
	id      string
	dto     *storageItem.DTO
}

func newItem(storage storage.Storage, dto *storageItem.DTO) item.Item {
	return &item_e{
		storage: storage,
		id:      dto.ID,
		dto:     dto,
	}
}

func (b *item_e) refresh() error {
	dto, err := b.storage.Item().Find(b.id)
	if err != nil {
		return err
	}

	b.dto = dto

	return nil
}

func (b *item_e) ID() string {
	return b.id
}

func (b *item_e) Data() (*item.Data, error) {
	if b.dto == nil {
		err := b.refresh()
		if err != nil {
			return nil, err
		}
	}

	data := &item.Data{
		Name: b.dto.Name,
	}

	return data, nil
}
