package box

import "errors"

func NewBox(repo Repository, params *CreateParams) (*Box, error) {
	if params == nil {
		return nil, errors.New("invalid params")
	}

	paramsRepo := &createParams{
		Name: params.Name,
		Type: params.Type,
	}

	if len(params.ParentID) > 0 {
		dto, err := repo.findById(params.ParentID)
		if err != nil {
			return nil, err
		}

		paramsRepo.Route = append(paramsRepo.Route, params.ParentID)

		for _, id := range dto.view().Route {
			paramsRepo.Route = append(paramsRepo.Route, id)
		}
	}

	dto, err := repo.create(paramsRepo)
	if err != nil {
		return nil, err
	}

	id := dto.view().ID
	node := &Box{ repo: repo, id: id, dto: dto }

	return node, nil
}

func FindBox(repo Repository, id string) (*Box, error) {
	dto, err := repo.findById(id)
	if err != nil {
		return nil, err
	}
	if dto == nil {
		return nil, nil
	}

	node := &Box{
		repo: repo,
		id: id,
		dto: dto,
	}

	return node, nil
}

func FindBoxes(repo Repository, filter *Filter) ([]*Box, error) {
	var nodes []*Box

	results, err := repo.findAll(filter)
	if err != nil {
		return nil, err
	}

	for _, dto := range results {
		id := dto.view().ID

		nodes = append(nodes, &Box{
			repo: repo,
			id: id,
			dto: dto,
		})
	}

	return nodes, nil
}