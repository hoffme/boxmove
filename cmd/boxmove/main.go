package main

import (
	"context"
	"log"
	"os"

	"github.com/hoffme/boxmove/clients"
	"github.com/hoffme/boxmove/server"
	"github.com/hoffme/boxmove/server/api"
	"github.com/hoffme/boxmove/storage"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	DbHost   := os.Getenv("DB_HOST")
	DbPort   := os.Getenv("DB_PORT")
	DbUser   := os.Getenv("DB_USER")
	DbPass   := os.Getenv("DB_PASS")
	DbName   := os.Getenv("DB_NAME")
	HttpAddr := os.Getenv("HTTP_ADDR")

	uri := "mongodb://"+ DbUser +":"+ DbPass +"@"+ DbHost +":"+ DbPort +"/"+ DbName
	ctx := context.TODO()

	conn, err := storage.GetConnection(uri, DbName, ctx)
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
