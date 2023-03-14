package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Conditions
	var currentYear = 2021

	if age := currentYear - 1998; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}

	// 2. Switch
	// Contoh 1
	var score = 8

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Contoh 2
	var score2 = 6

	switch {
	case score2 == 8:
		fmt.Println("perfect")
	case (score2 < 8) && (score2 > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}

	// Contoh 3
	var score3 = 6

	switch {
	case score3 == 8:
		fmt.Println("perfect")
	case (score3 < 8) && (score3 > 3):
		fmt.Println("not bad")
		fallthrough
	case score3 < 5:
		fmt.Println("It is ok, but please study harder")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}

	// Contoh 4
	var score4 = 10

	if score4 > 7 {
		switch score {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score4 == 5 {
			fmt.Println("not bad")
		} else if score == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score4 == 0 {
				fmt.Println("try harder!")
			}
		}
	}

	// 3. Loopings
	// Contoh 1
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}

	// Contoh 2
	var i = 0
	for i < 3 {
		fmt.Println("Angka", i)
		i++
	}

	// Contoh 3
	var i2 = 0

	for {
		fmt.Println("Angka", i2)

		i2++
		if i2 == 3 {
			break
		}
	}

	// Contoh 4
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	// Contoh 5
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// Contoh 6
	// outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke - ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				// break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}

	// 4. Array
	// Contoh 1
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var strings2 = [3]string{"Airell", "Nanda", "Mailo"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings2)

	// Contoh 2
	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"

	// Contoh 3
	var fruits2 = [3]string{"apple", "banana", "mango"}
	// Cara Pertama
	for i, v := range fruits2 {
		fmt.Printf("Index: %d, Value %s\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))

	// Cara Kedua
	for i := 0; i < len(fruits2); i++ {
		fmt.Printf("Index: %d, Value %s\n", i, fruits2[i])
	}

	// Contoh 4
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

	// 5. Slice
	// Contoh 1
	var fruits3 = []string{"apple", "banana", "mango"}
	_ = fruits3

	// Contoh 2
	var fruits4 = make([]string, 3)
	_ = fruits4
	fmt.Printf("%#v", fruits4)

	// Contoh 3
	var fruits5 = make([]string, 3)
	_ = fruits5

	fruits5[0] = "apple"
	fruits5[1] = "banana"
	fruits5[2] = "mango"

	fmt.Printf("%#v", fruits5)

	// Contoh 4
	var fruits6 = make([]string, 3)
	fruits6 = append(fruits6, "apple", "banana", "mango")
	fmt.Printf("%#v", fruits6)

	// Contoh 5
	var fruits7 = []string{"apple", "banana", "mango"}
	var fruits8 = []string{"durian", "pineapple", "starfruit"}
	fruits7 = append(fruits7, fruits8...)
	fmt.Printf("%#v", fruits7)

	// Contoh 6
	var fruits9 = []string{"apple", "banana", "mango"}
	var fruits10 = []string{"durian", "pineapple", "starfruit"}
	fruits11 := copy(fruits9, fruits10)
	fmt.Println("Fruits 9 =>", fruits9)
	fmt.Println("Fruits 10 =>", fruits10)
	fmt.Println("Copied elements =>", fruits11)

	// Contoh 7
	var fruits12 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var fruits13 = fruits12[1:4]
	fmt.Printf("%#v\n", fruits13)

	var fruits14 = fruits12[0:]
	fmt.Printf("%#v\n", fruits14)

	var fruits15 = fruits12[:3]
	fmt.Printf("%#v\n", fruits15)

	var fruits16 = fruits12[:] // sama dengan fruits12[:len(fruits12)]
	fmt.Printf("%#v\n", fruits16)

	// Contoh 8
	var fruits17 = []string{"apple", "banana", "mango", "durian", "pineapple"}
	fruits17 = append(fruits17[:3], "rambutan")
	fmt.Printf("%#v\n", fruits17)

	// Contoh 9
	var fruits18 = []string{"apple", "banana", "mango", "durian", "pineapple"}
	var fruits19 = fruits18[2:4]
	fruits19[0] = "rambutan"
	fmt.Println("fruits18 => ", fruits18)
	fmt.Println("fruits19 => ", fruits19)

	// Contoh 10
	var fruits20 = []string{"apple", "mango", "durian", "mango"}
	fmt.Println("Fruits20 cap:", cap(fruits20)) // 4
	fmt.Println("Fruits20 len:", len(fruits20)) // 4
	fmt.Println(strings.Repeat("#", 20))

	var fruits21 = fruits20[0:3]

	fmt.Println("Fruits21 cap:", cap(fruits21)) // 4
	fmt.Println("Fruits21 len:", len(fruits21)) // 3
	fmt.Println(strings.Repeat("#", 20))

	var fruits22 = fruits20[1:]

	fmt.Println("Fruits22 cap:", cap(fruits22)) // 3
	fmt.Println("Fruits22 len:", len(fruits22)) // 3

	// Contoh 11
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)

	// 6. Map
	// Contoh 1
	var person map[string]string // Deklarasi
	person = map[string]string{} // Inisialisasi

	person["name"] = "Airell"
	person["age"] = "23"
	person["address"] = "Jalan Sudirman"

	fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])

	// Contoh 2
	var person2 = map[string]string{
		"name":    "Airell",
		"age":     "23",
		"address": "Jalan Sudirman",
	}

	fmt.Println("name:", person2["name"])
	fmt.Println("age:", person2["age"])
	fmt.Println("address:", person2["address"])

	// Contoh 3
	var person3 = map[string]string{
		"name":    "Airell",
		"age":     "23",
		"address": "Jalan Sudirman",
	}

	for key, value := range person3 {
		fmt.Println(key, ":", value)
	}

	// Contoh 4
	var person4 = map[string]string{
		"name":    "Airell",
		"age":     "23",
		"address": "Jalan Sudirman",
	}

	fmt.Println("Before deleting:", person4)

	delete(person4, "age")

	fmt.Println("After deleting:", person4)

	// Contoh 5
	var person5 = map[string]string{
		"name":    "Airell",
		"age":     "23",
		"address": "Jalan Sudirman",
	}

	value, exist := person5["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value does'nt exist")
	}

	delete(person5, "age")

	value, exist = person5["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value does'nt exist")
	}

	// Contoh 6
	var people = []map[string]string{
		{"name": "Airell", "age": "23"},
		{"name": "Nanda", "age": "23"},
		{"name": "Mailo", "age": "15"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}

	// 7. Aliase
	// Contoh 1
	// deklarasi variable dengan tipe data uint8
	var a uint8 = 10
	var b byte // byte adala alias dari tipe data uint8

	b = a // no error, karena byte memiliki data yang sama dengan uint8
	_ = b

	// Contoh 2
	// Mendeklarasi tipe data alias bernama second
	// type nama_alias = nama_tipe_data
	type second = uint

	var hour second = 3600
	fmt.Printf("hour type: %T\n", hour) // => hour type: uint
}
