package move

import "time"

type DTO interface {
	View() *View
}

type Storage interface {
	FindById(id string) (DTO, error)
	FindAll(filter *DTOFilterParams) ([]DTO, error)
	Create(params *DTOCreateParams) (DTO, error)
}

type DTOFilterParams struct {
	ID       []string
	FromID   []string
	ToID     []string
	CountMin uint64
	CountMax uint64
	DateMin  time.Time
	DateMax  time.Time
}

type DTOCreateParams struct {
	FromID string
	ToID   string
	Date   time.Time
	Count  uint64
}