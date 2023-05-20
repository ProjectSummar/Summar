package main

import "log"

func main() {
	db, err := NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	apiServer := NewAPIServer("3001", db)
	apiServer.Run()
}
