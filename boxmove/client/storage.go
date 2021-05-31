package client

type DTO interface {
	ID()   string
	View() *View
}

type Storage interface {
	FindOne(id string) (DTO, error)
	FindAll(params *Filter) ([]DTO, error)
	Create(params *CreateParams) (DTO, error)
	Update(dto DTO, params *UpdateParams) error
	Delete(dto DTO) error
	Restore(dto DTO) error
	Remove(dto DTO) error
}
