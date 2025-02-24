package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initDB()
	router := mux.NewRouter()
	router.HandleFunc("/books", getBooksHandler)
	router.HandleFunc("/addBook", addBookHandler)
	router.HandleFunc("/borrowBook", borrowBookHandler)
	router.HandleFunc("/createBorrower", createBorrowerHandler)
	log.Println("app is running")
	log.Fatal(http.ListenAndServe(":8080", router))
}
