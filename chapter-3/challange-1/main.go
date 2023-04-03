// Buatlah sebuah service untuk melakukan POST data
// dalam format JSON dan secara acak setiap 15 detik dengan
// angka random antara 1-100 untuk valuewater dan wind.
// Gunakan url post berikut untuk menjalankan simulasi post request :

// "https://jsonplaceholder.typicode.com/posts"

// Kemudian tampilkan pada terminal hasil postnya. Selain itu kalian harus
// menentukan status water dan wind tersebut. Dengan ketentuan:

// - jika water dibawah 5 maka status aman
// - jika water antara 6 - 8 maka status siaga
// - jika water diatas 8 maka status bahaya
// - jika wind dibawah 6 maka status aman
// - jika wind antara 7 - 15 maka status siaga
// - jika wind diatas 15 maka status bahaya
// - value water dalam satuan meter
// - value wind dalam satuan meter per detik
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Post struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func getWaterStatus(water int) string {
	if water < 5 {
		return "aman"
	} else if water >= 6 && water <= 8 {
		return "siaga"
	} else {
		return "bahaya"
	}
}

func getWindStatus(wind int) string {
	if wind < 6 {
		return "aman"
	} else if wind >= 7 && wind <= 15 {
		return "siaga"
	} else {
		return "bahaya"
	}
}

func main() {
	for {
		// Mengirimkan POST request
		postBody := map[string]int{
			"water": rand.Intn(100) + 1,
			"wind":  rand.Intn(100) + 1,
		}

		postBodyBytes, err := json.MarshalIndent(postBody, "", "    ")
		if err != nil {
			fmt.Println("Error marshaling post body:", err)
			continue
		}

		fmt.Println(string(postBodyBytes))

		resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(postBodyBytes))
		if err != nil {
			fmt.Println("Error sending POST request:", err)
			continue
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			continue
		}

		var response Post
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println("Error unmarshaling response:", err)
			continue
		}

		output := fmt.Sprintf("status water : %s\n", getWaterStatus(postBody["water"]))
		output += fmt.Sprintf("status wind : %s\n", getWindStatus(postBody["wind"]))

		fmt.Println(output)

		time.Sleep(15 * time.Second)
	}
}
