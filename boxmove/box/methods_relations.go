package box

import "errors"

type TreeBox struct {
	Box      *Box       `json:"box"`
	Children []*TreeBox `json:"children"`
}

func (b *Box) Parent() (*Box, error) {
	view := b.View()
	if view == nil {
		return nil, errors.New("invalid box")
	}
	if len(view.Route) == 0 {
		return nil, nil
	}

	parentId := view.Route[0]

	dto, err := b.storage.FindOne(b.client, b.id)
	if err != nil {
		return nil, err
	}
	if dto == nil {
		return nil, nil
	}

	parent := &Box{
		storage: b.storage,
		id:      parentId,
		dto:     dto,
		client:  b.client,
	}

	return parent, nil
}

func (b *Box) HasAncestor(ancestor *Box) (bool, error) {
	view := b.View()
	if view == nil {
		return false, errors.New("invalid box")
	}

	for _, id := range view.Route {
		if ancestor.id == id{
			return true, nil
		}
	}

	return false, nil
}

func (b *Box) Ancestors() ([]*Box, error) {
	view := b.View()
	if view == nil {
		return nil, errors.New("invalid box")
	}

	result, err := b.storage.FindAll(b.client, &Filter{ ID: view.Route })
	if err != nil {
		return nil, err
	}

	var ancestors []*Box

	for _, dto := range result {
		box := &Box{
			storage: b.storage,
			id:      dto.View().ID,
			dto:     dto,
			client:  b.client,
		}

		ancestors = append(ancestors, box)
	}

	return ancestors, nil
}

func (b *Box) Children() ([]*Box, error) {
	var children []*Box

	results, err := b.storage.FindAll(b.client, &Filter{ ParentID: &b.id })
	if err != nil {
		return nil, err
	}

	for _, dto := range results {
		box := &Box{
			storage: b.storage,
			id:      dto.View().ID,
			dto:     dto,
			client:  b.client,
		}

		children = append(children, box)
	}

	return children, nil
}

func (b *Box) Decedents() ([]*Box, error) {
	results, err := b.storage.FindAll(b.client, &Filter{ AncestorID: b.id })
	if err != nil {
		return nil, err
	}

	var decedents []*Box

	for _, dto := range results {
		node := &Box{
			storage: b.storage,
			id:      dto.View().ID,
			dto:     dto,
			client:  b.client,
		}

		decedents = append(decedents, node)
	}

	return decedents, nil
}

func (b *Box) Tree() (*TreeBox, error) {
	results, err := b.storage.FindAll(b.client, &Filter{ AncestorID: b.id })
	if err != nil {
		return nil, err
	}

	root := &TreeBox{ Box: b }
	boxes := map[string]*TreeBox{ b.id: root }

	for _, dto := range results {
		box := &Box{
			storage: b.storage,
			id:      dto.View().ID,
			dto:     dto,
			client:  b.client,
		}

		boxes[box.id] = &TreeBox{ Box: box }
	}

	for _, treeBox := range boxes {
		view := treeBox.Box.View()
		if view == nil {
			return nil, errors.New("invalid view")
		}
		if len(view.Route) > 0 {
			continue
		}

		parentId := view.Route[0]

		parent := boxes[parentId]
		if parent == nil {
			return nil, errors.New("invalid results")
		}

		parent.Children = append(parent.Children, treeBox)
	}

	return root, nil
}
