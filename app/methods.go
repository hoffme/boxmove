package app

import (
	"errors"
	"log"

	"github.com/hoffme/boxmove/boxmove/box"
	"github.com/hoffme/boxmove/boxmove/move"
)

func (s *Service) NewMove(client string, params *move.CreateParams) (*move.Move, error) {
	extremes, err := s.Boxes.FindAll(client, &box.Filter{ ID: []string{ params.From, params.To } })
	if err != nil {
		return nil, err
	}

	if len(extremes) != 2 {
		return nil, errors.New("invalid extremes from move")
	}

	boxFrom := extremes[0]
	boxTo := extremes[1]
	if boxFrom.View().ID != params.From {
		boxTo = extremes[1]
		boxFrom = extremes[0]
	}

	if boxFrom.View().ID == boxTo.View().ID {
		return nil, errors.New("from box is same to box")
	}

	movement, err := s.Moves.New(client, params)
	if err != nil {
		return nil, err
	}

	errBoxFrom := boxFrom.AddActive(params.Active, int64(params.Count) * -1)
	if errBoxFrom != nil {
		log.Println(errBoxFrom)
	}

	errBoxTo := boxTo.AddActive(params.Active, int64(params.Count))
	if errBoxTo != nil {
		log.Println(errBoxTo)
	}

	return movement, nil
}

