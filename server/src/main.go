package main

import (
	"log"
	"os"
	"summar/server/handlers"
	"summar/server/routes"
	"summar/server/stores"
)

func main() {
	store, err := stores.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("ENV") == "dev" {
		store.Clear()
		log.Println("Cleared db")
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	handlers := handlers.NewHandlers(store)

	server := routes.NewServer()
	server.MountHandlers(handlers)

	address := os.Getenv("ADDRESS")
	server.Run(address)
}
