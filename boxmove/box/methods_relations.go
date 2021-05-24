package box

type TreeBox struct {
	Box      *Box       `json:"box"`
	Children []*TreeBox `json:"children"`
}

func (b *Box) Parent() (*Box, error) {
	view := b.View()
	if len(view.Route) == 0 {
		return nil, nil
	}

	parentId := view.Route[0]

	dto, err := b.repo.FindById(parentId)
	if err != nil {
		return nil, err
	}
	if dto == nil {
		return nil, nil
	}

	parent := &Box{
		repo: b.repo,
		id:   parentId,
		dto:  dto,
	}

	return parent, nil
}

func (b *Box) Ancestors() ([]*Box, error) {
	result, err := b.repo.FindAll(&DTOFilterParams{ ID: b.View().Route })
	if err != nil {
		return nil, err
	}

	var ancestors []*Box

	for _, dto := range result {
		box := &Box{
			repo: b.repo,
			id:   dto.View().ID,
			dto:  dto,
		}

		ancestors = append(ancestors, box)
	}

	return ancestors, nil
}

func (b *Box) HasAncestor(ancestor *Box) (bool, error) {
	for _, id := range b.View().Route {
		if ancestor.id == id{
			return true, nil
		}
	}
	return false, nil
}

func (b *Box) Children() ([]*Box, error) {
	var children []*Box

	results, err := b.repo.FindAll(&DTOFilterParams{ ParentID: &b.id })
	if err != nil {
		return nil, err
	}

	for _, dto := range results {
		box := &Box{
			repo: b.repo,
			id:   dto.View().ID,
			dto:  dto,
		}

		children = append(children, box)
	}

	return children, nil
}

func (b *Box) Decedents() ([]*Box, error) {
	results, err := b.repo.FindAll(&DTOFilterParams{ AncestorID: &b.id })
	if err != nil {
		return nil, err
	}

	var decedents []*Box

	for _, dto := range results {
		node := &Box{
			repo: b.repo,
			id:   dto.View().ID,
			dto:  dto,
		}

		decedents = append(decedents, node)
	}

	return decedents, nil
}

func (b *Box) Tree() (*TreeBox, error) {
	results, err := b.repo.FindAll(&DTOFilterParams{ AncestorID: &b.id })
	if err != nil {
		return nil, err
	}

	root := &TreeBox{ Box: b }

	boxes := map[string]*TreeBox{}
	for _, dto := range results {
		box := &Box{
			repo: b.repo,
			id:   dto.View().ID,
			dto:  dto,
		}
		boxes[box.id] = &TreeBox{ Box: box}
	}

	for _, treeNode := range boxes {
		parentId := treeNode.Box.View().Route[0]

		var parent *TreeBox
		if parentId == b.id {
			parent = root
		} else {
			parent = boxes[parentId]
		}

		parent.Children = append(parent.Children, treeNode)
	}

	return root, nil
}
