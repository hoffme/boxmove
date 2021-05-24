package main

import (
	"context"
	"log"
	"os"

	"github.com/hoffme/boxmove/management"
	"github.com/hoffme/boxmove/server"
	"github.com/hoffme/boxmove/server/api"
	"github.com/hoffme/boxmove/storage"
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	DbUri    := getEnv("DB_URI", "mongodb://localhost:27017")
	DbName   := getEnv("DB_NAME", "mongo")
	HttpAddr := getEnv("HTTP_ADDR", ":3000")

	ctx := context.TODO()

	conn, err := storage.GetConnection(DbUri, DbName, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	man := management.New(conn)

	srv := server.New(HttpAddr, api.CreateRouter(man))

	err = srv.Run()
	if err != nil {
		log.Fatal(err)
	}
}
