package boxmove

import (
	"errors"

	"github.com/hoffme/boxmove/boxmove/box"
	"github.com/hoffme/boxmove/boxmove/move"
	"github.com/hoffme/boxmove/storage"
)

type Management struct {
	boxRepo box.Repository
	movRepo move.Repository
}

func New(connection *storage.Connection, key string) (*Management, error) {
	boxRepo, err := box.NewMongoRepository(connection, "boxes", key)
	if err != nil {
		return nil, errors.New("not created box repository: " + err.Error())
	}

	movRepo, err := move.NewMongoRepository(connection, "moves", key)
	if err != nil {
		return nil, errors.New("not created move repository: " + err.Error())
	}

	f := &Management{
		boxRepo: boxRepo,
		movRepo: movRepo,
	}

	return f, nil
}
