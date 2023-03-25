package main

import (
	"challange-5/routers"
)

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)

	// Simple Rest API Buat Simple Rest API dengan item book dengan method untuk :
	// A. Get All Book
	// B. Get Book by id
	// C. Add Book
	// D. Update Book
	// E. Delete Book
}
