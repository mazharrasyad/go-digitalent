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
}
