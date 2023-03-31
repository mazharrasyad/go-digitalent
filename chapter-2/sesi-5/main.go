package main

import (
	"chapter-2/sesi-5/routers"
	"encoding/json"
	"fmt"
	"net/url"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type User struct {
	FullName string `json:"Name"`
	Age      int
}

const PORT = ":3000"

func main() {
	// 1. Decoding & Parsin JSON Data
	var jsonString = `
		{
			"full_name": "Airell Jordan",
			"email": "airell@mail.com",
			"age": 23
		}
	`

	var jsonString2 = `[
			{
				"full_name": "Airell Jordan",
				"email": "airell@mail.com",
				"age": 23
			},
			{
				"full_name": "Ananda RHP",
				"email": "ananda@mail.com",
				"age": 23
			}
		]		
	`

	// Contoh 1
	var result Employee

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name", result.FullName)
	fmt.Println("email", result.Email)
	fmt.Println("age", result.Age)

	// Contoh 2
	var result2 map[string]interface{}

	var err2 = json.Unmarshal([]byte(jsonString), &result2)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	fmt.Println("full_name", result2["full_name"])
	fmt.Println("email", result2["email"])
	fmt.Println("age", result2["age"])

	// Contoh 3
	var temp interface{}

	var err3 = json.Unmarshal([]byte(jsonString), &temp)
	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}

	var result3 = temp.(map[string]interface{})

	fmt.Println("full_name", result3["full_name"])
	fmt.Println("email", result3["email"])
	fmt.Println("age", result3["age"])

	// Contoh 4
	var result4 []Employee

	var err4 = json.Unmarshal([]byte(jsonString2), &result4)
	if err4 != nil {
		fmt.Println(err4.Error())
		return
	}

	for i, v := range result4 {
		fmt.Printf("Index %d: %v+\n", i+1, v)
	}

	// Contoh 5
	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}

	var jsonData, err5 = json.Marshal(object)

	if err5 != nil {
		fmt.Println(err5.Error())
		return
	}

	var jsonString3 = string(jsonData)
	fmt.Println(jsonString3)

	// 2. URL Parsing
	var urlString = "http://developer.com:80/hello?name=airell&age=23"
	var u, e = url.Parse(urlString)

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)
	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name: %s, age: %s\n", name, age)

	// 3. Swagger
	routers.Router().Run(PORT)
}
