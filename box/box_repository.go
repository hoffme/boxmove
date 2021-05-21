package box

type Repository interface {
	findById(id string) (dto, error)
	findAll(filter *Filter) ([]dto, error)
	create(fields *createParams) (dto, error)
	update(dto dto, params *updateParams) error
	restore(dto dto) error
	delete(dto dto) error
	remove(dto dto) error
}

type dto interface {
	view() *View
}

type createParams struct {
	Route []string
	Name  string
	Type  string
}

type updateParams struct {
	Name  *string
	Route *[]string
}