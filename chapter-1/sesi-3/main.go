package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// 1. Function
	// Contoh 1
	greet("Airell", 23)

	// Contoh 2
	greet2("Airell", "jalan sudirman")

	// Contoh 3
	var names = []string{"Airell", "Jordan"}
	var printMsg = greet3("Heii", names)

	fmt.Println(printMsg)

	// Contoh 4
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Println("Area:", area)
	fmt.Println("Circumference:", circumference)
	// 2. a
	// 3. b
	// 4. c
}

// Contoh 1
func greet(name string, age int8) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old", name, age)
}

// Contoh 2
func greet2(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in", address)
}

// Contoh 3
func greet3(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")
	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

// Contoh 4
func calculate(d float64) (float64, float64) {
	// Menghitung luas
	var area float64 = math.Pi * math.Pow(d/2, 2)

	// Menghitung keliling
	var circumference = math.Pi * d

	return area, circumference
}

// Contoh 5
