package box

import "errors"

func NewBox(repo Storage, params *CreateParams) (*Box, error) {
	if params == nil {
		return nil, errors.New("invalid params")
	}

	paramsRepo := &DTOCreateParams{
		Name: params.Name,
		Type: params.Type,
	}

	if len(params.ParentID) > 0 {
		dto, err := repo.FindById(params.ParentID)
		if err != nil {
			return nil, err
		}

		paramsRepo.Route = append(paramsRepo.Route, params.ParentID)

		for _, id := range dto.View().Route {
			paramsRepo.Route = append(paramsRepo.Route, id)
		}
	}

	dto, err := repo.Create(paramsRepo)
	if err != nil {
		return nil, err
	}

	id := dto.View().ID
	node := &Box{ repo: repo, id: id, dto: dto }

	return node, nil
}

func FindBox(repo Storage, id string) (*Box, error) {
	dto, err := repo.FindById(id)
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

func FindBoxes(repo Storage, filter *Filter) ([]*Box, error) {
	results, err := repo.FindAll(&DTOFilterParams{
		ID: filter.ID,
		ParentID: filter.ParentID,
		AncestorID: filter.AncestorID,
		Name: filter.Name,
		Type: filter.Type,
		Deleted: filter.Deleted,
	})
	if err != nil {
		return nil, err
	}

	var boxes []*Box

	for _, dto := range results {
		box := &Box{
			repo: repo,
			id: dto.View().ID,
			dto: dto,
		}

		boxes = append(boxes, box)
	}

	return boxes, nil
}