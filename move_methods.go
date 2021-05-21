package boxmove

import (
	"github.com/hoffme/boxmove/box"
	"github.com/hoffme/boxmove/move"
)

/* Move Methods */

func (m *Management) NewMove(params *move.CreateParams) (*move.Move, error) {
	return move.NewMove(m.movRepo, params)
}

func (m *Management) FindMove(id string) (*move.Move, error) {
	return move.FindMove(m.movRepo, id)
}

func (m *Management) FindMoves(filter *move.Filter) ([]*move.Move, error) {
	return move.FindMoves(m.movRepo, filter)
}

/* Extremes Move */

func (m *Management) MoveExtremes(mv *move.Move) (*box.Box, *box.Box, error) {
	ids := []string{mv.View().FromID, mv.View().ToID }

	nodes, err := box.FindBoxes(m.boxRepo, &box.Filter{ ID: ids })
	if err != nil {
		return nil, nil, err
	}
	if len(nodes) == 0 {
		return nil, nil, nil
	}

	var from *box.Box
	var to *box.Box

	if nodes[0].View().ID == mv.View().FromID {
		from = nodes[0]
		if len(nodes) == 2 {
			to = nodes[1]
		}
	} else {
		from = nodes[1]
		if len(nodes) == 2 {
			to = nodes[0]
		}
	}


	return from, to, nil
}

func (m *Management) MoveFromBox(mv *move.Move) (*box.Box, error) {
	return box.FindBox(m.boxRepo, mv.View().FromID)
}

func (m *Management) MoveToBox(mv *move.Move) (*box.Box, error) {
	return box.FindBox(m.boxRepo, mv.View().ToID)
}

