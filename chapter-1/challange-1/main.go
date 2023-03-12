package main

import "fmt"

func main() {
	// Buatlah sebuah program go dengan :

	// 1. menampilkan nilai i : 21 fmt.Printf("%T \n", i) // contoh : fmt.Printf("%v \n", i)
	// expected output : 21
	fmt.Println("1. menampilkan nilai i : 21")
	i := 21
	fmt.Printf("%v \n", i)

	// 2. menampilkan tipe data dari variabel i
	// expected output : int
	fmt.Println("\n2. menampilkan tipe data dari variabel i")
	fmt.Printf("%T \n", i)

	// 3. menampilkan tanda %
	// expected output : %
	fmt.Println("\n3. menampilkan tanda %")
	fmt.Printf("%%\n")

	// 4. menampilkan nilai boolean j : true
	// expected output : true
	fmt.Println("\n4. menampilkan nilai boolean j : true")
	j := true
	fmt.Println(j)

	// 5. menampilkan unicode russia : Я (ya) - RALAT
	// fmt.Println("\n5. menampilkan unicode russia : Я (ya)")
	// Я := "\u042F"
	// fmt.Println(Я)

	// 5. menampilkan nilai base2 dari i : 10101
	fmt.Println("\n5. menampilkan nilai base2 dari i : 10101")
	fmt.Printf("%b\n", 10101)

	// 6. menampilkan nilai base 10 : 21
	// expected output : 21
	fmt.Println("\n6. menampilkan nilai base 10 : 21")
	fmt.Printf("%d\n", 21)

	// 7. menampilkan nilai base 8 : 25
	// expected output : 25
	fmt.Println("\n7. menampilkan nilai base 8 : 25")
	fmt.Printf("%o\n", 25)

	// 8. menampilkan nilai base 16 : f
	// expected output : f
	fmt.Println("\n8. menampilkan nilai base 16 : f")
	fmt.Printf("%x\n", 'f')

	// 9. menampilkan nilai base 16 : F 13
	// expected output : F
	fmt.Println("\n9. menampilkan nilai base 16 : F")
	fmt.Printf("%X\n", "F")

	// 10. menampilkan unicode karakter Я : U+042F var k float64 = 123.456;
	// expected output : U+042F
	fmt.Println("\n10. menampilkan unicode karakter Я : U+042F var k float64 = 123.456")
	fmt.Printf("%U dan %.3f\n", 'Я', 123.456)

	// 11. menampilkan float : 123.456000
	// expected output : 123.456000
	fmt.Println("\n11. menampilkan float : 123.456000")
	fmt.Printf("%f\n", 123.456000)

	// 12. menampilkan float scientific : 1.234560E+02
	// expected output : 1.234560E+02
	fmt.Println("\n12. menampilkan float scientific : 1.234560E+02")
	fmt.Printf("%.6E\n", 1.234560e+02)
}
