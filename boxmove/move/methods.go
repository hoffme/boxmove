package move

func NewMove(repo Repository, params *CreateParams) (*Move, error) {
	dto, err := repo.create(params)
	if err != nil {
		return nil, err
	}

	id := dto.view().ID
	tr := &Move{ repo: repo, id: id, dto: dto }

	return tr, nil
}

func FindMove(repo Repository, id string) (*Move, error) {
	dto, err := repo.findById(id)
	if err != nil {
		return nil, err
	}
	if dto == nil {
		return nil, nil
	}

	move := &Move{ repo: repo, id: id, dto: dto }

	return move, nil
}

func FindMoves(repo Repository, filter *Filter) ([]*Move, error) {
	var moves []*Move

	results, err := repo.findAll(filter)
	if err != nil {
		return nil, err
	}

	for _, dto := range results {
		id := dto.view().ID

		moves = append(moves, &Move{
			repo: repo,
			id: id,
			dto: dto,
		})
	}

	return moves, nil
}