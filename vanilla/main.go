package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("GET /ping", handelePing)

	log.Println("Server listining...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}

type response struct {
	Msg string `json:"message"`
}

func handelePing(w http.ResponseWriter, r *http.Request) {

	res := response{Msg: "pong"}

	json.NewEncoder(w).Encode(res)

	w.WriteHeader(http.StatusOK)

	log.Println("Request rescives")
}
