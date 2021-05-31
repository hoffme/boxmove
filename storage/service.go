package storage

import (
	"context"
	"log"

	"github.com/hoffme/boxmove/storage/connections"

	"github.com/hoffme/boxmove/storage/models/active"
	"github.com/hoffme/boxmove/storage/models/box"
	"github.com/hoffme/boxmove/storage/models/client"
	"github.com/hoffme/boxmove/storage/models/move"

	"github.com/hoffme/boxmove/utils"
)

type Service struct {
	mongoConnection *connections.MongoConnection

	ClientStorage *client.Storage
	ActiveStorage *active.Storage
	BoxStorage 	  *box.Storage
	MoveStorage   *move.Storage
}

func NewService() (*Service, error) {
	service := &Service{}

	err := service.setConnections()
	if err != nil {
		return nil, err
	}

	err = service.setStorages()
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (s *Service) setConnections() error {
	uri := utils.GetEnv("DB_URI", "mongodb://localhost:27017")
	db  := utils.GetEnv("DB_NAME", "active")
	ctx := context.Background()

	mongo, err := connections.MongoDB(uri, db, ctx)
	if err != nil {
		return err
	}

	s.mongoConnection = mongo

	return nil
}

func (s *Service) setStorages() error {
	clientStorage, err := client.New(s.mongoConnection, "clients")
	if err != nil {
		return err
	}

	activeStorage, err := active.New(s.mongoConnection, "actives")
	if err != nil {
		return err
	}

	boxesStorage, err := box.New(s.mongoConnection, "boxes")
	if err != nil {
		return err
	}

	movesStorage, err := move.New(s.mongoConnection, "moves")
	if err != nil {
		return err
	}

	s.ClientStorage = clientStorage
	s.ActiveStorage = activeStorage
	s.BoxStorage    = boxesStorage
	s.MoveStorage   = movesStorage

	return nil
}

func (s *Service) Close() {
	err := s.mongoConnection.Close()
	if err != nil {
		log.Fatal(err)
	}
}