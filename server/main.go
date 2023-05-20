package main

import "log"

func main() {
	db, err := NewPostgresDB()
	if err != nil {
		log.Fatal("Cannot connect to postgres DB")
	}

	if err := db.Init(); err != nil {
		log.Fatal("Cannot initialise postgres DB")
	}

	apiServer := NewAPIServer("3001", db)
	apiServer.Run()
}
