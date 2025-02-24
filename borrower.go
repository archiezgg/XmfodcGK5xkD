package main

import "gorm.io/gorm"

type borrower struct {
	gorm.Model
	bookID    uint
	firstName string
	lastName  string
}
