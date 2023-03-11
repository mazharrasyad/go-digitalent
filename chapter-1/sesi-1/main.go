package main

import "fmt"

func main() {
	// 1. fmt
	fmt.Println("<isi-pesan>")
	fmt.Println("Hello World")
	fmt.Println("Hello World", "Message", "From", "Go")

	fmt.Println("Airell Jordan")
	fmt.Println("Airell", "Jordan")

	fmt.Print("Airell Jordan\n")
	fmt.Print("Airell", " Jordan\n")
	fmt.Print("Airell", " ", "Jordan\n")

	// 2. Komentar
	// komentar kode
	// menghasilkan pesan Hello World

	/*
		komentar kode
		menghasilkan pesan Hello World
	*/

	fmt.Println("Hello World")

	// fmt.Println("Baris ini tidak akan dieksekusi")

	// 3. Variables
	// Contoh 1
	var name string = "Airell"
	var age int = 23

	fmt.Println("Ini adalah namanya ==>", name)
	fmt.Println("Ini adalah umurnya ==>", age)

	// Contoh 2
	var name2 string
	var age2 int = 23

	name2 = "Airell"
	age2 = 25

	fmt.Println("Ini adalah namanya2 ==>", name2)
	fmt.Println("Ini adalah umurnya2 ==>", age2)

	// Contoh 3
	var name3 string
	var age3 int

	// Error
	// name3 = 30
	// age3 = "Airell"

	fmt.Println("Ini adalah namanya3 ==>", name3)
	fmt.Println("Ini adalah umurnya3 ==>", age3)

	// Contoh 4
	var name4 = "Airell"
	var age4 = 23

	fmt.Printf("%T, %T", name4, age4)

	// Contoh 5
	name5 := "Airell"
	age5 := 23

	fmt.Printf("%T, %T", name5, age5)

	// Contoh 6
	var student1, student2, student3 string = "satu", "dua", "tiga"

	var first, second, third int

	first, second, third = 1, 2, 3

	fmt.Println(student1, student2, student3)

	fmt.Println(first, second, third)

	// Contoh 7
	var name7, age7, address7 = "Airell", 23, "Jalan sudirman"

	first7, second7, third7 := "1", 2, "3"

	fmt.Println(name7, age7, address7)

	fmt.Println(first7, second7, third7)

	// Contoh 8
	var firstVariable string

	var name8, age8, address8 = "Airell", 23, "Jalan sudirman"

	_, _, _, _ = firstVariable, name8, age8, address8

	// Contoh 9
	first9, second9 := 1, "2"

	fmt.Printf("Tipe data variable first9 adalah %T \n", first9)
	fmt.Printf("Tipe data variable second9 adalah %T \n", second9)

	// Contoh 10
	var nama10 = "Airell"
	var age10 = 23
	var address10 = "Jalan sudirman"
	fmt.Printf("Halo nama ku %s, umurku aku adalah %d, dan aku tinggal di %s", nama10, age10, address10)

	// 4. Data Type
	// Contoh 1
	var first41 = 89
	var second41 = 12

	fmt.Printf("tipe data first41 : %T\n", first41)
	fmt.Printf("tipe data second41 : %T\n", second41)

	// Contoh 2
	var first42 uint8 = 89
	var second42 int8 = 12

	fmt.Printf("tipe data first42 : %T\n", first42)
	fmt.Printf("tipe data second42 : %T\n", second42)

	// Contoh 3
	var decimalNumber float32 = 3.63

	fmt.Printf("decimal Number: %f\n", decimalNumber)
	fmt.Printf("decimal Number: %.3f\n", decimalNumber)

	// 5. Bool
	var condition bool = true
	fmt.Printf("is it permitted? %t \n", condition)

	// 6. String
	// Contoh 1
	var message string = "Halo"
	fmt.Print(message)

	// Contoh 2
	var message2 = `Selamat datang di "Hactiv8". 
Salam kenal :).
	Mari belajar "Scalable Web Service With Go".`

	fmt.Println(message2)

	// 7. Constants
	// Contoh 1
	const full_name string = "Airell Jordan"

	fmt.Printf("Hello %s", full_name)

	// Contoh 2
	// Error
	// const full_name2 string

	// fmt.Println(full_name2)

	// 8. IOTA
	// Contoh 1
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3) // => 0 1 2

	// Contoh 2
	const (
		c11 = iota + 2 // => 2
		c22            // => 3
		c33            // => 4
	)

	const (
		c44 = iota - 2 // => -2
		c55            // => -1
		c66            // => 0
	)

	const (
		c77 = iota * 2 // => 0
		c88            // => 2
		c99            // => 4
	)

	const (
		c100 = iota / 2 // => 0
		c101            // => 0
		c102            // => 1
	)

	fmt.Println(c11, c22, c33)
	fmt.Println(c44, c55, c66)
	fmt.Println(c77, c88, c99)
	fmt.Println(c100, c101, c102)

	// 9. Operators
	// Contoh 1
	var value = 2 + 2*3
	fmt.Println(value)

	// Contoh 2
	var value2 = (2 + 2) * 3
	fmt.Println(value2)

	// Contoh 3
	var firstCondition bool = 2 < 3
	var secondCondition bool = "joey" == "Joey"
	var thirdCondition bool = 10 != 2.3
	var fourthCondition bool = 11 <= 11

	fmt.Println("first condition:", firstCondition)
	fmt.Println("second condition:", secondCondition)
	fmt.Println("third condition:", thirdCondition)
	fmt.Println("fourth condition:", fourthCondition)

	// Contoh 4
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t(%t) \n", wrongReverse)
}
