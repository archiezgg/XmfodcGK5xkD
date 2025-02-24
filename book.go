package main

import (
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BorrowerID uint
	Title      string
}

func getAllBooks() ([]Book, error) {
	var books []Book
	result := database.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func getBookByID(bookID uint) (Book, error) {
	var b Book
	result := database.First(&b, bookID)
	if result.Error != nil {
		return Book{}, result.Error
	}
	return b, nil
}

func addBook(title string) error {
	book := Book{
		Title:      title,
		BorrowerID: 0, //as a first implementation, creation happens with borrowerID as 0
	}

	result := database.Create(&book)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("new book added: title: %v, borrower: -", title)
	return nil
}

func borrowBook(bookID uint, borrowerID uint) error {
	book, err := getBookByID(bookID)
	if err != nil {
		return err
	}

	book.BorrowerID = borrowerID

	result := database.Save(&book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
