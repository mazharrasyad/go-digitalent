package main

import (
	"fmt"
	"strings"
)

func main() {
	kalimat := "selamat malam"
	hitungKalimat(kalimat)
}

func hitungKalimat(kalimat string) {
	jumlahKata := strings.Split(kalimat, " ")
	var jumlahHuruf = make(map[string]int)

	for _, kata := range jumlahKata {
		for _, huruf := range kata {
			fmt.Println(string(huruf))
			jumlahHuruf[string(huruf)]++
		}
		fmt.Println()
	}

	fmt.Println(jumlahHuruf)
}
