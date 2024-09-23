package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	connStr := "user=koffee dbname=mockdata password=koffee host=103.166.182.81 port=5432 sslmode=disable" // Adjust this string
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
}
