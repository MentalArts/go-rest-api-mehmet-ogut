package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("GET /ping", handelePing)

	log.Println("Server listining...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func handelePing(w http.ResponseWriter, r *http.Request) {

	log.Println("Request rescives")
}
