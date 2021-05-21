package main

import (
	"context"
	"log"

	"github.com/hoffme/boxmove/storage"
)

func main() {
	DBURL  := "mongodb://127.0.0.1:27017/"
	DBNAME := "boxmove"
	//HTTPAddr      := "localhost:5000"

	ctx := context.TODO()

	conn, err := storage.GetConnection(DBURL, DBNAME, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}
