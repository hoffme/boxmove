package box

import "time"

type Box struct {
	repo Repository
	id   string
	dto  dto
}

type View struct {
	ID 		  string   	 `json:"id"`
	Route	  []string 	 `json:"route"`
	Name 	  string   	 `json:"name"`
	Type      string   	 `json:"type"`
	CreatedAt time.Time	 `json:"created_at"`
	UpdatedAt time.Time	 `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type TreeBox struct {
	Box      *Box       `json:"box"`
	Children []*TreeBox `json:"children"`
}

func (b *Box) View() *View {
	return b.dto.view()
}

func (b *Box) Refresh() error {
	dto, err := b.repo.findById(b.id)
	if err != nil {
		return err
	}

	b.dto = dto

	return nil
}

func (b *Box) EqualTo(node *Box) bool {
	return b.id == node.id
}

func (b *Box) Parent() (*Box, error) {
	view := b.View()
	if len(view.Route) == 0 {
		return nil, nil
	}

	parentId := view.Route[0]

	dto, err := b.repo.findById(parentId)
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
	var ancestors []*Box

	result, err := b.repo.findAll(&Filter{ ID: b.View().Route })
	if err != nil {
		return nil, err
	}

	for _, dto := range result {
		box := &Box{
			repo: b.repo,
			id:   dto.view().ID,
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

func (b *Box) update(params *updateParams) error {
	err := b.repo.update(b.dto, params)
	if err != nil {
		return err
	}

	return nil
}

func (b *Box) SetParent(box *Box) error {
	newRoute := &[]string{ box.id }

	for _, ancestorID := range box.View().Route {
		*newRoute = append(*newRoute, ancestorID)
	}

	err := b.update(&updateParams{ Route: newRoute })
	if err != nil {
		return err
	}

	return nil
}

func (b *Box) SetName(name string) error {
	return b.update(&updateParams{ Name: &name })
}

func (b *Box) Children() ([]*Box, error) {
	var children []*Box

	results, err := b.repo.findAll(&Filter{ ParentID: &b.id })
	if err != nil {
		return nil, err
	}

	for _, dto := range results {
		id := dto.view().ID

		children = append(children, &Box{
			repo: b.repo,
			id:   id,
			dto:  dto,
		})
	}

	return children, nil
}

func (b *Box) Decedents() ([]*Box, error) {
	results, err := b.repo.findAll(&Filter{ AncestorID: &b.id })
	if err != nil {
		return nil, err
	}

	var decedents []*Box

	for _, dto := range results {
		node := &Box{
			repo: b.repo,
			id:   dto.view().ID,
			dto:  dto,
		}
		decedents = append(decedents, node)
	}

	return decedents, nil
}

func (b *Box) Tree() (*TreeBox, error) {
	results, err := b.repo.findAll(&Filter{ AncestorID: &b.id })
	if err != nil {
		return nil, err
	}

	root := &TreeBox{ Box: b }

	boxes := map[string]*TreeBox{}
	for _, dto := range results {
		box := &Box{
			repo: b.repo,
			id:   dto.view().ID,
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

func (b *Box) Delete() error {
	err := b.repo.delete(b.dto)
	if err != nil {
		return err
	}

	return nil
}

func (b *Box) Remove() error {
	err := b.repo.remove(b.dto)
	if err != nil {
		return err
	}

	return nil
}