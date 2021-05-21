package boxmove

import (
	box2 "github.com/hoffme/boxmove/boxmove/box"
	move2 "github.com/hoffme/boxmove/boxmove/move"
)

/* Box Methods */

func (m *Management) NewBox(params *box2.CreateParams) (*box2.Box, error) {
	return box2.NewBox(m.boxRepo, params)
}

func (m *Management) FindBox(id string) (*box2.Box, error) {
	return box2.FindBox(m.boxRepo, id)
}

func (m *Management) FindBoxes(filter *box2.Filter) ([]*box2.Box, error) {
	return box2.FindBoxes(m.boxRepo, filter)
}

/* Stats Box */

func (m *Management) BoxCount(b *box2.Box, options *BoxCountsOptions) (uint64, error) {
	_, _, total, err := m.BoxCountStats(b, options)
	return total, err
}

func (m *Management) BoxCountStats(b *box2.Box, options *BoxCountsOptions) (uint64, uint64, uint64, error) {
	moves, err := m.BoxMoves(b, &BoxMovesOptions{
		Ingress: true,
		Egress: true,
		DateMin: options.DateMin,
		DateMax: options.DateMax,
		Deleted: options.Deleted,
	})
	if err != nil {
		return 0, 0, 0, err
	}

	var ingress uint64
	var egress uint64

	for _, mov := range moves {
		count := mov.View().Count

		if mov.View().FromID == b.View().ID {
			egress += count
		} else {
			ingress += count
		}
	}

	return ingress, egress, ingress - egress, nil
}

func (m *Management) BoxCountIngress(b *box2.Box, options *BoxCountsOptions) (uint64, error) {
	transfers, err := m.BoxMoves(b, &BoxMovesOptions{
		Ingress: true,
		DateMin: options.DateMin,
		DateMax: options.DateMax,
		Deleted: options.Deleted,
	})
	if err != nil {
		return 0, err
	}

	var ingress uint64

	for _, trn := range transfers {
		count := trn.View().Count
		ingress += count
	}

	return ingress, nil
}

func (m *Management) BoxCountEgress(b *box2.Box, options *BoxCountsOptions) (uint64, error) {
	transfers, err := m.BoxMoves(b, &BoxMovesOptions{
		Egress: true,
		DateMin: options.DateMin,
		DateMax: options.DateMax,
		Deleted: options.Deleted,
	})
	if err != nil {
		return 0, err
	}

	var egress uint64

	for _, trn := range transfers {
		count := trn.View().Count
		egress += count
	}

	return egress, nil
}

/* Movements Box */

func (m *Management) BoxMoves(b *box2.Box, options *BoxMovesOptions) ([]*move2.Move, error) {
	decedents, err := b.Decedents()
	if err != nil {
		return nil, err
	}
	if decedents == nil || len(decedents) == 0 {
		return []*move2.Move{}, nil
	}

	var ids []string
	for _, boxDecedent := range decedents {
		ids = append(ids, boxDecedent.View().ID)
	}

	filter := &move2.Filter{
		CountMin: options.CountMin,
		CountMax: options.CountMax,
		DateMin: options.DateMin,
		DateMax: options.DateMax,
	}

	if options.Egress {
		filter.FromID = ids
	}
	if options.Ingress {
		filter.ToID = ids
	}

	return move2.FindMoves(m.movRepo, filter)
}
