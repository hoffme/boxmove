package client

import (
	"github.com/hoffme/boxmove/boxmove/box"
	"github.com/hoffme/boxmove/boxmove/move"
)

/* Move Methods */

func (c *Client) NewMove(params *move.CreateParams) (*move.Move, error) {
	return move.NewMove(c.storageMove, params)
}

func (c *Client) FindMove(id string) (*move.Move, error) {
	return move.FindMove(c.storageMove, id)
}

func (c *Client) FindMoves(filter *move.Filter) ([]*move.Move, error) {
	return move.FindMoves(c.storageMove, filter)
}

/* Extremes Move */

func (c *Client) MoveExtremes(mv *move.Move) (*box.Box, *box.Box, error) {
	ids := []string{mv.View().FromID, mv.View().ToID }

	nodes, err := box.FindBoxes(c.storageBox, &box.Filter{ ID: ids })
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

func (c *Client) MoveFromBox(mv *move.Move) (*box.Box, error) {
	return box.FindBox(c.storageBox, mv.View().FromID)
}

func (c *Client) MoveToBox(mv *move.Move) (*box.Box, error) {
	return box.FindBox(c.storageBox, mv.View().ToID)
}

