package client

import "time"

type Storage interface {
	New(params *DTOCreateParams) (DTO, error)
	Get(id string) (DTO, error)
	Delete(dto DTO) error
}

type DTO interface {
	View() *View
}

type DTOCreateParams struct {
	Name 	  string
	CreatedAt time.Time
}