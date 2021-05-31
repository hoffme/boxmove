package active

import "errors"

type Active struct {
	storage Storage
	dto     DTO
	id      string
	client  string
}

func (a *Active) View() *View {
	if a.dto == nil {
		return nil
	}
	return a.dto.View()
}

func (a *Active) Refresh() error {
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

func (a *Active) Update(params *UpdateParams) error {
	if params == nil {
		return errors.New("invalid params")
	}

	err := a.storage.Update(a.dto, params)
	if err != nil {
		return err
	}

	return nil
}

func (a *Active) Delete() error {
	err := a.storage.Delete(a.dto)
	if err != nil {
		return err
	}

	return nil
}

func (a *Active) Restore() error {
	err := a.storage.Restore(a.dto)
	if err != nil {
		return err
	}

	return nil
}

func (a *Active) Remove() error {
	err := a.storage.Remove(a.dto)
	if err != nil {
		return err
	}

	return nil
}
