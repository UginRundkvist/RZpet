package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// тут будем регистрировать хендлеры

	log.Println("test server started on :8081")

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}
