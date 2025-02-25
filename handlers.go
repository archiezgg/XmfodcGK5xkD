package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
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

func borrowBookHandler(w http.ResponseWriter, r *http.Request) {
	type requestedBody struct {
		BookID     string `json:"bookId"`
		BorrowerID string `json:"borrowerId"`
	}

	var rb requestedBody
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	bookId, err := strconv.ParseUint(rb.BookID, 10, 64)
	if err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	borrowerId, err := strconv.ParseUint(rb.BorrowerID, 10, 64)
	if err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	if err := borrowBook(uint(bookId), uint(borrowerId)); err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	msg := fmt.Sprintf("book with ID: %v successfully borrowed by user with ID: %v", rb.BookID, rb.BorrowerID)
	writeMessage(w, msg)
}

func createBorrowerHandler(w http.ResponseWriter, r *http.Request) {
	type requestedBody struct {
		Username string `json:"username"`
	}

	var rb requestedBody
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if err := createBorrower(rb.Username); err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	writeMessage(w, "borrower created")
}

func getBorrowerHandler(w http.ResponseWriter, r *http.Request) {
	type requestedBody struct {
		BorrowerID string `json:"borrowerId"`
	}

	var rb requestedBody
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	borrowerId, err := strconv.ParseUint(rb.BorrowerID, 10, 64)
	if err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	borrower, err := getBorrowerByID(uint(borrowerId))
	if err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	msg := fmt.Sprintf("the borrower details: Username: %v, Borrowed Books: %v", borrower.UserName, borrower.Books)
	writeMessage(w, msg)
}

func getBorrowedBooksHandler(w http.ResponseWriter, r *http.Request) {
	type requestedBody struct {
		BorrowerID string `json:"borrowerId"`
	}

	var rb requestedBody
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	borrowerId, err := strconv.ParseUint(rb.BorrowerID, 10, 64)
	if err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	borrowedBooks, err := getBorrowedBooksByBorrowerID(uint(borrowerId))
	if err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	msg := fmt.Sprintf("the borrowed books by %v are: %v", borrowerId, borrowedBooks)
	writeMessage(w, msg)
}
