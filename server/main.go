package main

import "log"

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if ENV == "DEV" {
		store.Clear()
		log.Println("Cleared db")
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	apiServer := NewAPIServer("3001", store)
	apiServer.Run()
}
