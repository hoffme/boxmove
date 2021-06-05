package box

import "errors"

type Box struct {
	storage Storage
	dto     DTO
	id      string
	client  string
}

func (b *Box) Id() string {
	return b.id
}

func (b *Box) View() *View {
	if b.dto == nil {
		return nil
	}
	return b.dto.View()
}

func (b *Box) Refresh() error {
	dto, err := b.storage.FindOne(b.id, b.client)
	if err != nil {
		return err
	}
	if dto == nil {
		return errors.New("not found")
	}

	b.dto = dto

	return nil
}

func (b *Box) AddActive(active string, count int64) error {
	view := b.View()
	if view == nil {
		return errors.New("invalid box")
	}

	newActives := map[string]int64{}
	for id, count := range view.Actives {
		newActives[id] = count
	}

	newActives[active] += count

	if newActives[active] == 0 {
		delete(newActives, active)
	}

	return b.Update(&UpdateParams{ Actives: &newActives })
}

func (b *Box) SetActive(active string, count int64) error {
	view := b.View()
	if view == nil {
		return errors.New("invalid box")
	}

	newActives := map[string]int64{}
	for id, count := range view.Actives {
		newActives[id] = count
	}

	if count == 0 {
		delete(newActives, active)
	} else {
		newActives[active] = count
	}

	return b.Update(&UpdateParams{ Actives: &newActives })
}

func (b *Box) Update(params *UpdateParams) error {
	if params == nil {
		return errors.New("invalid params")
	}

	err := b.storage.Update(b.dto, params)
	if err != nil {
		return err
	}

	return nil
}

func (b *Box) Delete() error {
	err := b.storage.Delete(b.dto)
	if err != nil {
		return err
	}

	return nil
}

func (b *Box) Restore() error {
	err := b.storage.Restore(b.dto)
	if err != nil {
		return err
	}

	return nil
}

func (b *Box) Remove() error {
	err := b.storage.Remove(b.dto)
	if err != nil {
		return err
	}

	return nil
}
