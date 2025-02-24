package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func logErrorAndSendHTTPError(w http.ResponseWriter, err error, httpStatusCode int) {
	log.Println(err)
	errorMsg := fmt.Sprintf("{\"error\": \"%v\"}", err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write([]byte(errorMsg))
}

func writeMessage(w http.ResponseWriter, msg string) {
	finalMessage := fmt.Sprintf("{\"message\": \"%s\"}", msg)
	w.Write([]byte(finalMessage))
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := getAllBooks()
	if err != nil {
		log.Println("can't retrieve Books")
	}

	for _, v := range books {
		w.Write([]byte(v.Title + "\n"))
	}
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
	type requestedBody struct {
		Title string `json:"title"`
	}

	var rb requestedBody
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if err := addBook(rb.Title); err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	writeMessage(w, "book created successfully")
}
