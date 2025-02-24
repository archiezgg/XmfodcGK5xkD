package main

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	BorrowerID uint
	Author     string
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
