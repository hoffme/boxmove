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

	box := &Box{
		repo: repo,
		id: id,
		dto: dto,
	}

	return box, nil
}

func FindBoxes(repo Repository, filter *Filter) ([]*Box, error) {
	results, err := repo.findAll(filter)
	if err != nil {
		return nil, err
	}

	var boxes []*Box

	for _, dto := range results {
		box := &Box{
			repo: repo,
			id: dto.view().ID,
			dto: dto,
		}

		boxes = append(boxes, box)
	}

	return boxes, nil
}