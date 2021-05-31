package storage

import (
	"context"
	"log"

	"github.com/hoffme/boxmove/utils"
)

type Service struct {
	Mongo *MongoConnection
}

func NewService() (*Service, error) {
	uri := utils.GetEnv("DB_URI", "mongodb://localhost:27017")
	db  := utils.GetEnv("DB_NAME", "mongo")
	ctx := context.Background()

	mongo, err := getMongo(uri, db, ctx)
	if err != nil {
		return nil, err
	}

	service := &Service{ Mongo: mongo }

	return service, nil
}

func (s *Service) Close() {
	err := s.Mongo.Close()
	if err != nil {
		log.Fatal(err)
	}
}