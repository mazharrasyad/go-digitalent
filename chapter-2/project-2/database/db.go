package database

import (
	"fmt"
	"log"
	"project-2/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbPort   = 5432
	dbname   = "learning-gorm"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user =%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to datbase :", err)
	}

	db.Debug().AutoMigrate(models.Book{})
}

func GetDB() *gorm.DB {
	return db
}
