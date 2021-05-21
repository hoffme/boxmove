package boxmove

import (
	"github.com/hoffme/boxmove/box"
	"github.com/hoffme/boxmove/move"
)

/* Box Methods */

func (m *Management) NewBox(params *box.CreateParams) (*box.Box, error) {
	return box.NewBox(m.boxRepo, params)
}

func (m *Management) FindBox(id string) (*box.Box, error) {
	return box.FindBox(m.boxRepo, id)
}

func (m *Management) FindBoxes(filter *box.Filter) ([]*box.Box, error) {
	return box.FindBoxes(m.boxRepo, filter)
}

/* Stats Box */

func (m *Management) BoxCount(b *box.Box, options *BoxCountsOptions) (uint64, error) {
	_, _, total, err := m.BoxCountStats(b, options)
	return total, err
}

func (m *Management) BoxCountStats(b *box.Box, options *BoxCountsOptions) (uint64, uint64, uint64, error) {
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

func (m *Management) BoxCountIngress(b *box.Box, options *BoxCountsOptions) (uint64, error) {
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

func (m *Management) BoxCountEgress(b *box.Box, options *BoxCountsOptions) (uint64, error) {
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

func (m *Management) BoxMoves(b *box.Box, options *BoxMovesOptions) ([]*move.Move, error) {
	decedents, err := b.Decedents()
	if err != nil {
		return nil, err
	}
	if decedents == nil || len(decedents) == 0 {
		return []*move.Move{}, nil
	}

	var ids []string
	for _, boxDecedent := range decedents {
		ids = append(ids, boxDecedent.View().ID)
	}

	filter := &move.Filter{
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

	return move.FindMoves(m.movRepo, filter)
}
