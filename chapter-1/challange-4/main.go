package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var listBiodata = []Biodata{
	{"Muhammad", "Jl. Depok No. 1", "Web Developer", "Ingin integrasi golang dengan web"},
	{"Azhar", "Jl. Depok No. 2", "Mobile Developer", "Ingin integrasi golang dengan mobile"},
	{"Rasyad", "Jl. Depok No. 3", "System Analyst", "Ingin memahami proses kerja golang"},
}

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage: go run main.go <absen>")
		return
	}

	absen := args[1]

	biodata, found := findBiodata(absen)

	if found {
		fmt.Println("Nama:", biodata.Nama)
		fmt.Println("Alamat:", biodata.Alamat)
		fmt.Println("Pekerjaan:", biodata.Pekerjaan)
		fmt.Println("Alasan:", biodata.Alasan)
	} else {
		fmt.Println("Biodata tidak ada")
	}
}

func findBiodata(absen string) (Biodata, bool) {
	index := parseAbsen(absen)

	if index < 0 || index >= len(listBiodata) {
		return Biodata{}, false
	}

	biodata := listBiodata[index]

	return biodata, true
}

func parseAbsen(absen string) int {
	var index int
	_, err := fmt.Sscan(absen, &index)

	if err != nil {
		return -1
	}

	return index - 1
}
