package main

import (
	"log"
	"summar/server/constants"
	"summar/server/handlers"
	"summar/server/postgres"
	"summar/server/routes"
)

func main() {
	store, err := postgres.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if constants.ENV == "DEV" {
		store.Clear()
		log.Println("Cleared db")
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	handlers := handlers.NewHandlers(store)

	apiServer := routes.NewRouter("3001", handlers)
	apiServer.Run()
}
