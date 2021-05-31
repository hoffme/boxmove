package active

type DTO interface {
	ID()   string
	View() *View
}

type Storage interface {
	FindOne(client, id string) (DTO, error)
	FindAll(client string, params *Filter) ([]DTO, error)
	Create(client string, params *CreateParams) (DTO, error)
	Update(dto DTO, params *UpdateParams) error
	Delete(dto DTO) error
	Restore(dto DTO) error
	Remove(dto DTO) error
}
