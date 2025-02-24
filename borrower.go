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
