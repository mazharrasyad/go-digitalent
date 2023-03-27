package main

import (
	"project-2/routers"
)

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
