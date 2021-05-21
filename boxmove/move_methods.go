package boxmove

import (
	box2 "github.com/hoffme/boxmove/boxmove/box"
	move2 "github.com/hoffme/boxmove/boxmove/move"
)

/* Move Methods */

func (m *Management) NewMove(params *move2.CreateParams) (*move2.Move, error) {
	return move2.NewMove(m.movRepo, params)
}

func (m *Management) FindMove(id string) (*move2.Move, error) {
	return move2.FindMove(m.movRepo, id)
}

func (m *Management) FindMoves(filter *move2.Filter) ([]*move2.Move, error) {
	return move2.FindMoves(m.movRepo, filter)
}

/* Extremes Move */

func (m *Management) MoveExtremes(mv *move2.Move) (*box2.Box, *box2.Box, error) {
	ids := []string{mv.View().FromID, mv.View().ToID }

	nodes, err := box2.FindBoxes(m.boxRepo, &box2.Filter{ ID: ids })
	if err != nil {
		return nil, nil, err
	}
	if len(nodes) == 0 {
		return nil, nil, nil
	}

	var from *box2.Box
	var to *box2.Box

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

func (m *Management) MoveFromBox(mv *move2.Move) (*box2.Box, error) {
	return box2.FindBox(m.boxRepo, mv.View().FromID)
}

func (m *Management) MoveToBox(mv *move2.Move) (*box2.Box, error) {
	return box2.FindBox(m.boxRepo, mv.View().ToID)
}

