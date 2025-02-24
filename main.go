package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	log.Println("app is running")
	log.Fatal(http.ListenAndServe(":8080", router))
}
