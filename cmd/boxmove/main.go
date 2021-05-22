package main

import (
	"context"
	"log"
	"os"

	"github.com/hoffme/boxmove/clients"
	"github.com/hoffme/boxmove/server"
	"github.com/hoffme/boxmove/server/api"
	"github.com/hoffme/boxmove/storage"
)

func main() {
	DbUri    := os.Getenv("DB_URI")
	DbName   := os.Getenv("DB_NAME")
	HttpAddr := os.Getenv("HTTP_ADDR")

	ctx := context.TODO()

	conn, err := storage.GetConnection(DbUri, DbName, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cls := clients.New(conn)

	srv := server.New(HttpAddr, api.CreateRouter(cls))

	err = srv.Run()
	if err != nil {
		log.Fatal(err)
	}
}
