package main

import (
	"fmt"
)

func main() {
	// Buatlah sebuah program go dengan implementasi perulangan for dan kombinasi if-else dengan expected output :
	i := 0
	for j := 0; j <= 10; j++ {
		if i == 4 {
			j = -1
		}

		if i <= 4 {
			fmt.Println("Nilai i = ", i)
			i++
		} else if j == 5 {
			const kalimat = "CAIAPBO"
			for index, runeValue := range kalimat {
				fmt.Printf("character %U starts at byte position %d\n", runeValue, index*2)
			}
		} else {
			fmt.Println("Nilai j = ", j)
		}
	}
}
