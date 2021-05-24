package box

type Storage interface {
	FindById(id string) (DTO, error)
	FindAll(params *DTOFilterParams) ([]DTO, error)
	Create(params *DTOCreateParams) (DTO, error)
	Update(dto DTO, params *DTOUpdateParams) error
	Restore(dto DTO) error
	Delete(dto DTO) error
	Remove(dto DTO) error
}

type DTO interface {
	View() *View
}

type DTOFilterParams struct {
	ID         []string
	ParentID   *string
	AncestorID *string
	Name       string
	Type       string
	Deleted    bool
}

type DTOCreateParams struct {
	Route []string
	Name  string
	Type  string
}

type DTOUpdateParams struct {
	Name  *string
	Route *[]string
}