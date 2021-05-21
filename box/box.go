package box

import (
	"errors"
	"time"
)

type Box struct {
	repo Repository
	id   string
	dto  dto
}

type View struct {
	ID 		  string   	 `json:"id"`
	Route	  []string 	 `json:"route"`
	Name 	  string   	 `json:"name"`
	Type      string   	 `json:"type"`
	CreatedAt time.Time	 `json:"created_at"`
	UpdatedAt time.Time	 `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (b *Box) View() *View {
	if b.dto == nil {
		return nil
	}

	return b.dto.view()
}

func (b *Box) Refresh() error {
	dto, err := b.repo.findById(b.id)
	if err != nil {
		return err
	}

	b.dto = dto

	return nil
}

func (b *Box) EqualTo(node *Box) bool {
	return b.id == node.id
}

func (b *Box) SetName(name string) error {
	err := b.repo.update(b.dto, &updateParams{ Name: &name })
	if err != nil {
		return err
	}

	return nil
}

func (b *Box) SetParent(box *Box) error {
	newRoute := &[]string{ box.id }

	for _, ancestorID := range box.View().Route {
		*newRoute = append(*newRoute, ancestorID)
	}

	err := b.repo.update(b.dto, &updateParams{ Route: newRoute })
	if err != nil {
		return err
	}

	return nil
}

func (b *Box) Delete() error {
	err := b.repo.delete(b.dto)
	if err != nil {
		return err
	}

	return nil
}

func (b *Box) Remove() error {
	err := b.Refresh()
	if err != nil {
		return err
	}

	if b.View().DeletedAt == nil {
		return errors.New("the box is not deleted")
	}

	err = b.repo.remove(b.dto)
	if err != nil {
		return err
	}

	b.dto = nil

	return nil
}