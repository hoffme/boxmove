package client

import "errors"

type Client struct {
	storage Storage
	dto     DTO
	id      string
}

func (a *Client) Id() string {
	return a.id
}

func (a *Client) View() *View {
	if a.dto == nil {
		return nil
	}
	return a.dto.View()
}

func (a *Client) Refresh() error {
	dto, err := a.storage.FindOne(a.id)
	if err != nil {
		return err
	}
	if dto == nil {
		return errors.New("not found")
	}

	a.dto = dto

	return nil
}

func (a *Client) Update(params *UpdateParams) error {
	if params == nil {
		return errors.New("invalid params")
	}

	err := a.storage.Update(a.dto, params)
	if err != nil {
		return err
	}

	return nil
}

func (a *Client) Delete() error {
	err := a.storage.Delete(a.dto)
	if err != nil {
		return err
	}

	return nil
}

func (a *Client) Restore() error {
	err := a.storage.Restore(a.dto)
	if err != nil {
		return err
	}

	return nil
}

func (a *Client) Remove() error {
	err := a.storage.Remove(a.dto)
	if err != nil {
		return err
	}

	return nil
}
