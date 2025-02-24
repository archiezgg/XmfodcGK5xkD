package main

import (
	"log"
	"net/http"
)

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := getAllBooks()
	if err != nil {
		log.Println("can't retrieve Books")
	}

	for _, v := range books {
		w.Write([]byte(v.Title + "\n"))
	}
}
