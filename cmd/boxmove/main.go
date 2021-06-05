package main

import (
	"log"

	"github.com/hoffme/boxmove/app"
	_interface "github.com/hoffme/boxmove/interface"
	"github.com/hoffme/boxmove/storage"
)

func main() {
	storageService, err := storage.NewService()
	if err != nil {
		log.Fatal(err)
	}
	defer storageService.Close()

	appService, err := app.NewService(storageService)
	if err != nil {
		log.Fatal(err)
	}

	interfaceService, err := _interface.NewService(appService)
	if err != nil {
		log.Fatal(err)
	}

	interfaceService.Start()
}
