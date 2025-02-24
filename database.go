package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database   *gorm.DB
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PWD")
	dbName     = os.Getenv("DB_NAME")
)

func initDB() {
	dbInfo := fmt.Sprintf("host=%v port=%v user=%v "+
		"password=%v dbname=%v sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	var err error
	database, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		try := 1
		for try <= 6 && err != nil {
			log.Printf("establishing connection to the database... %d\nExiting after 5 tries.", try)
			time.Sleep(2 * time.Second)
			database, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
			try++
			if try == 6 {
				panic(err)
			}
		}
	}

	database.AutoMigrate(&Book{})
	database.AutoMigrate(&Borrower{})
	log.Println("connected to DB")
}
