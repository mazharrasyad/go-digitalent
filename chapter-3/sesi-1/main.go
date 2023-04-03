package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	h8HelperRand "github.com/novalagung/gubrak/v2"
)

func main() {
	// 1. Vendoring
	fmt.Println(h8HelperRand.RandomInt(10, 20))

	// 2. HTTP Request
	// GET
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	sb := string(body)
	log.Println(sb)

	// POST
	data := map[string]interface{}{
		"title":  "Airell",
		"body":   "Jordan",
		"userId": 1,
	}

	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	res2, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res2.Body.Close()

	body2, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body2))

	// 3. HTTP Basic AUTH
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("server started at localhost:9000")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGET(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
