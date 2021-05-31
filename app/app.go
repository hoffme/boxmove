package app

import (
	"github.com/hoffme/boxmove/storage"

	"github.com/hoffme/boxmove/boxmove/active"
	"github.com/hoffme/boxmove/boxmove/box"
	"github.com/hoffme/boxmove/boxmove/client"
	"github.com/hoffme/boxmove/boxmove/move"

	activeMongo "github.com/hoffme/boxmove/boxmove/active/mongo"
	boxMongo "github.com/hoffme/boxmove/boxmove/box/mongo"
	clientMongo "github.com/hoffme/boxmove/boxmove/client/mongo"
	moveMongo "github.com/hoffme/boxmove/boxmove/move/mongo"
)

type Service struct {
	Clients *client.Store
	Actives *active.Store
	Boxes 	*box.Store
	Moves 	*move.Store
}

func NewAppService(storage *storage.Service) (*Service, error) {
	clientStorage, err := clientMongo.New(storage.Mongo, "clients")
	if err != nil {
		return nil, err
	}

	activeStorage, err := activeMongo.New(storage.Mongo, "actives")
	if err != nil {
		return nil, err
	}

	boxesStorage, err := boxMongo.New(storage.Mongo, "boxes")
	if err != nil {
		return nil, err
	}

	movesStorage, err := moveMongo.New(storage.Mongo, "moves")
	if err != nil {
		return nil, err
	}

	service := &Service{
		Clients: &client.Store{ Storage: clientStorage },
		Actives: &active.Store{ Storage: activeStorage },
		Boxes:   &box.Store{ Storage: boxesStorage },
		Moves:   &move.Store{ Storage: movesStorage },
	}

	return service, nil
}
