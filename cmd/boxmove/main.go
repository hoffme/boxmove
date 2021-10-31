package main

import (
	"fmt"
	"log"

	"github.com/hoffme/boxmove/pkg/core/box"
	"github.com/hoffme/boxmove/pkg/core/item"
	"github.com/hoffme/boxmove/pkg/core/move"

	boxInt "github.com/hoffme/boxmove/internal/core/box"
	itemInt "github.com/hoffme/boxmove/internal/core/item"
	moveInt "github.com/hoffme/boxmove/internal/core/move"

	"github.com/hoffme/boxmove/internal/storage/mongo"
)

func main() {
	storage, err := mongo.New(&mongo.Settings{
		Uri:                "mongodb://127.0.0.1:27017/",
		Database:           "boxmove",
		BoxCollectionName:  "boxes",
		ItemCollectionName: "items",
		MoveCollectionName: "moves",
	})
	if err != nil {
		log.Println("create storage")
		log.Fatal(err)
	}
	defer storage.Close()

	boxStore := boxInt.NewStore(storage)
	moveStore := moveInt.NewStore(storage)
	itemStore := itemInt.NewStore(storage)

	finanzas, err := boxStore.Create(&box.StoreCreateParams{
		Name: "Finanzas",
	})
	if err != nil {
		log.Println("create finanzas box")
		log.Fatal(err)
	}

	caja_principal, err := boxStore.Create(&box.StoreCreateParams{
		Name:   "Caja Principal",
		Parent: finanzas.ID(),
	})
	if err != nil {
		log.Println("create caja principal box")
		log.Fatal(err)
	}

	caja_diaria, err := boxStore.Create(&box.StoreCreateParams{
		Name:   "Caja Diaria",
		Parent: finanzas.ID(),
	})
	if err != nil {
		log.Println("create caja diaria box")
		log.Fatal(err)
	}

	ars, err := itemStore.Create(&item.StoreCreateParams{
		Name: "ARS",
	})
	if err != nil {
		log.Println("create ars item")
		log.Fatal(err)
	}

	move, err := moveStore.Create(&move.StoreCreateParams{
		FromID:   caja_diaria.ID(),
		ToID:     caja_principal.ID(),
		ItemID:   ars.ID(),
		Quantity: 1500000,
	})
	if err != nil {
		log.Println("create move")
		log.Fatal(err)
	}

	fmt.Print(move)
}
