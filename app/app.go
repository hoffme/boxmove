package app

import (
	"github.com/hoffme/boxmove/storage"

	"github.com/hoffme/boxmove/boxmove/active"
	"github.com/hoffme/boxmove/boxmove/box"
	"github.com/hoffme/boxmove/boxmove/client"
	"github.com/hoffme/boxmove/boxmove/move"
)

type Service struct {
	Clients *client.Store
	Actives *active.Store
	Boxes 	*box.Store
	Moves 	*move.Store
}

func NewAppService(storage *storage.Service) (*Service, error) {
	service := &Service{
		Clients: &client.Store{ Storage: storage.ClientStorage },
		Actives: &active.Store{ Storage: storage.ActiveStorage },
		Boxes:   &box.Store{ Storage: storage.BoxStorage },
		Moves:   &move.Store{ Storage: storage.MoveStorage },
	}

	return service, nil
}
