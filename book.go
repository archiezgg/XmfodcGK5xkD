package main

import "gorm.io/gorm"

type book struct {
	gorm.Model
	borrowerID uint
	author     string
	title      string
}
