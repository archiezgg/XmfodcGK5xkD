package main

import (
	"log"

	"gorm.io/gorm"
)

type Borrower struct {
	gorm.Model
	Books    []Book
	UserName string
}

func createBorrower(userName string) error {
	borrower := Borrower{
		UserName: userName,
		Books:    nil,
	}

	result := database.Create(&borrower)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("new borrower added: username: %v", userName)
	return nil
}

func getBorrowerByID(borrowerID uint) (Borrower, error) {
	var b Borrower
	result := database.First(&b, borrowerID)

	if result.Error != nil {
		return Borrower{}, result.Error
	}

	borrowedBooks, err := getBorrowedBooksByBorrowerID(borrowerID)
	if err != nil {
		return Borrower{}, err
	}

	b.Books = borrowedBooks
	return b, nil
}

func getBorrowedBooksByBorrowerID(borrowerID uint) ([]Book, error) {
	var books []Book
	result := database.Where("borrower_id = ?", borrowerID).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
