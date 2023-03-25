package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// 1. FizzBuzz
	var n int // inisialisasi variabel n
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n) // mengambil nilai n dari input user

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// Defer & Exit, Error Handling
	// Contoh 1
	defer fmt.Println("defer function starts to execute")
	fmt.Println("Hai everyone")
	fmt.Println("Welcome back to Go learning center")

	// Contoh 2
	callDeferFunc()
	fmt.Println("Hai everyone!!")

	// Contoh 3
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before Exiting")
	// os.Exit(1) // Keluar Program

	// 3. Error
	// Contoh 1
	var number int
	var err error

	number, err = strconv.Atoi("123GH")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	number, err = strconv.Atoi("123")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	// Contoh 2
	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valid)
	}
}

// Defer
func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}

// Error
func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}

	return "Valid password", nil
}
