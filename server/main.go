package main

import "log"

func main() {
	db, err := NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	if ENV == "DEV" {
		db.Clear()
		log.Println("Cleared db")
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	apiServer := NewAPIServer("3001", db)
	apiServer.Run()
}
