package main

import (
	"chapter-3/challange-3/database"
	"chapter-3/challange-3/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
