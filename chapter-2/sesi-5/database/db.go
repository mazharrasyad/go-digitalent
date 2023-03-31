package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST     = "localhost"
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_PORT     = 5432
	DB_NAME     = "learning-gorm"
	DEBUG_MODE  = true
)

var db *gorm.DB

func New() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
