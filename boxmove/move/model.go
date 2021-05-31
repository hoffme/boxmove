package move

import "errors"

type Move struct {
	storage Storage
	dto     DTO
	id      string
	client  string
}

func (a *Move) View() *View {
	if a.dto == nil {
		return nil
	}
	return a.dto.View()
}

func (a *Move) Refresh() error {
	dto, err := a.storage.FindOne(a.id, a.client)
	if err != nil {
		return err
	}
	if dto == nil {
		return errors.New("not found")
	}

	a.dto = dto

	return nil
}
