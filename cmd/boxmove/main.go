package main

import (
	"log"

	"github.com/hoffme/boxmove/internal/app"
	"github.com/hoffme/boxmove/internal/storage/mongo"
)

func main() {
	storage, err := mongo.New("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = app.SetStorage(storage)
	if err != nil {
		log.Fatal(err)
	}
}
