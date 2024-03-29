package main

import (
	"chapter-3/sesi-2/database"
	"chapter-3/sesi-2/router"
)

func main() {
	// mux := http.NewServeMux()

	// endpoint := http.HandlerFunc(greet)

	// mux.Handle("/", middleware1(middleware2(endpoint)))

	// fmt.Println("Listening to port 8080")

	// err := http.ListenAndServe(":3000", mux)

	// log.Fatal(err)

	// JSON
	database.StartDB()
	r := router.StartApp()
	r.Run(":8081")
}

// func greet(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello World!!!"))
// }

// func middleware1(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("middleware pertama")
// 		next.ServeHTTP(w, r)
// 	})
// }

// func middleware2(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("middleware kedua")
// 		next.ServeHTTP(w, r)
// 	})
// }
