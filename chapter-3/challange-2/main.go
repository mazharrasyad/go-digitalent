// Membuat API dengan middleware authentication, authorization multi level user, dan
// authorization product by user id

// Detail Challenge :Buatlah Rest API product (create, read, update, delete) dengan fitur login dan register, 
// serta memiliki 3 fitur middleware antara lain :
// -Authentication
// -Authorization multi level user
// -Authorization access product by id
// Notes : buatlah authentication dengan JWT token golang, 
// lalu gunakan token tersebut untuk setiap hit Rest API product.

package main

import (
	"chapter-3/challange-2/database"
	"chapter-3/challange-2/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}