package move

func NewMove(repo Storage, params *CreateParams) (*Move, error) {
	dto, err := repo.Create(&DTOCreateParams{
		FromID: params.FromID,
		ToID: params.ToID,
		Date: params.Date,
		Count: params.Count,
	})
	if err != nil {
		return nil, err
	}

	id := dto.View().ID
	tr := &Move{ repo: repo, id: id, dto: dto }

	return tr, nil
}

func FindMove(repo Storage, id string) (*Move, error) {
	dto, err := repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if dto == nil {
		return nil, nil
	}

	move := &Move{ repo: repo, id: id, dto: dto }

	return move, nil
}

func FindMoves(repo Storage, filter *Filter) ([]*Move, error) {
	var moves []*Move

	results, err := repo.FindAll(&DTOFilterParams{
		ID: filter.ID,
		FromID: filter.FromID,
		ToID: filter.ToID,
		CountMin: filter.CountMin,
		CountMax: filter.CountMax,
		DateMin: filter.DateMin,
		DateMax: filter.DateMax,
	})
	if err != nil {
		return nil, err
	}

	for _, dto := range results {
		id := dto.View().ID

		moves = append(moves, &Move{
			repo: repo,
			id: id,
			dto: dto,
		})
	}

	return moves, nil
}