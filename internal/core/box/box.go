package box

import (
	"github.com/hoffme/boxmove/pkg/core/box"

	"github.com/hoffme/boxmove/pkg/storage"
	storageBox "github.com/hoffme/boxmove/pkg/storage/box"
)

type box_e struct {
	storage storage.Storage
	id      string
	dto     *storageBox.DTO
}

func newBox(storage storage.Storage, dto *storageBox.DTO) box.Box {
	return &box_e{
		storage: storage,
		id:      dto.ID,
		dto:     dto,
	}
}

func (b *box_e) refresh() error {
	dto, err := b.storage.Box().Find(b.id)
	if err != nil {
		return err
	}

	b.dto = dto

	return nil
}

func (b *box_e) ID() string {
	return b.id
}

func (b *box_e) Data() (*box.Data, error) {
	if b.dto == nil {
		err := b.refresh()
		if err != nil {
			return nil, err
		}
	}

	data := &box.Data{
		Name:  b.dto.Name,
		Route: b.dto.Route,
	}

	return data, nil
}
