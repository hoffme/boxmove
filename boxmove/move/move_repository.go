package move

type dto interface {
	view() *View
}

type Repository interface {
	findById(id string) (dto, error)
	findAll(filter *Filter) ([]dto, error)
	create(params *CreateParams) (dto, error)
}
