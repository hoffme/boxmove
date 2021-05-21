package main

import (
	"context"
	"github.com/hoffme/boxmove/clients"
	"github.com/hoffme/boxmove/server"
	"github.com/hoffme/boxmove/server/api"
	"log"

	"github.com/hoffme/boxmove/storage"
)

func main() {
	DBURL    := "mongodb://127.0.0.1:27017/"
	DBNAME   := "boxmove"
	HTTPADDR := "localhost:5000"

	ctx := context.TODO()

	conn, err := storage.GetConnection(DBURL, DBNAME, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cls := clients.New(conn)

	srv := server.New(HTTPADDR, api.CreateRouter(cls))

	err = srv.Run()
	if err != nil {
		log.Fatal(err)
	}
}
