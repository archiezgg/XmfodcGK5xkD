package main

import "gorm.io/gorm"

type Borrower struct {
	gorm.Model
	BookID    uint
	FirstName string
	LastName  string
}
