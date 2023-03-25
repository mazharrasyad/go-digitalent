package main

import (
	"challange-6/routers"
)

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
