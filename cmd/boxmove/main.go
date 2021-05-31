package main

import (
	"github.com/hoffme/boxmove/app"
	"github.com/hoffme/boxmove/server"
	"log"

	"github.com/hoffme/boxmove/storage"
)

func main() {
	storageService, err := storage.NewService()
	if err != nil {
		log.Fatal(err)
	}
	defer storageService.Close()

	appService, err := app.NewAppService(storageService)
	if err != nil {
		log.Fatal(err)
	}

	serverService, err := server.NewService(appService)
	if err != nil {
		log.Fatal(err)
	}

	serverService.Start()
}
