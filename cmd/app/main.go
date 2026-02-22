package main

import (
	"automation-service/internal/storage"
	"log"
)

func main() {
	dsn := "postgres://admin:admin@localhost:5432/automation?sslmode=disable"

	pool, err := storage.New(dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	log.Println("Connected to database successfully")

	// пока просто держим программу живой
	select {}
}
