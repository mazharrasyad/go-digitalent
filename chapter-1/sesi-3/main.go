package main

import (
	"fmt"
	"math"
	"reflect"
	"sesi-3/helpers"
	"strings"
)

// Closure
type isOddNum func(int) bool

// Struct
// Contoh 1
type Employee struct {
	name     string
	age      int
	division string
}

// Contoh 2
type Person struct {
	name string
	age  int
}

type EmployeeA struct {
	division string
	person   Person
}

// Method
// Contoh 1
func (p Person) Introduce(msg string) string {
	return fmt.Sprintf("%s My name is %s and I'm %d years old", msg, p.name, p.age)
}

// Contoh 2
func (p Person) ChangeName1() {
	p.name = "Mailo"
}

func (p *Person) ChangeName2() {
	p.name = "Mailo"
}

// Contoh 3
func (p *Person) Greet() {
	fmt.Println("Haii everyone")
}

// Reflect
// Contoh 1
type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama		:", reflectType.Field(i).Name)
		fmt.Println("tipe data	:", reflectType.Field(i).Type)
		fmt.Println("nilai		:", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

// Contoh 2
func (s *student) SetName(name string) {
	s.Name = name
}

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

	// Contoh 5
	studentLists := print("Airell", "Nanda", "Mailo", "Schannel", "Marco")
	fmt.Printf("%v", studentLists)

	// Contoh 6
	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := sum(numberLists...)

	fmt.Println("Result:", result)

	// Contoh 7
	profile("Airell", "pasta", "ayam geprek", "ikan roa", "sate padang")

	// 2. Closure
	// Contoh 1
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}

	var numbers = []int{4, 93, 77, 10, 52, 22, 34}

	fmt.Println(evenNumbers(numbers...))

	// Contoh 2
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("katak")

	fmt.Println(isPalindrome)

	// Contoh 3
	var studentLists2 = []string{"Airell", "Nanda", "Mailo", "Schannel", "Marca"}
	var find = findStudent(studentLists2)
	fmt.Println(find("airell"))

	// Contoh 4
	var numbers2 = []int{2, 5, 8, 10, 3, 99, 23}
	var find2 = findOddNumbers(numbers2, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers:", find2)

	// Contoh 5
	var numbers3 = []int{2, 5, 8, 10, 3, 99, 23}
	var find3 = findOddNumbers2(numbers3, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers:", find3)

	// 3. Pointer
	// Contoh 1
	var firstNumber *int
	var secondNumber *int
	_, _ = firstNumber, secondNumber

	// Contoh 2
	var firstNumber2 int = 4
	var secondNumber2 *int = &firstNumber2

	fmt.Println("firstNumber2 (value)	:", firstNumber2)
	fmt.Println("firstNumber2 (memori address)	:", &firstNumber2)

	fmt.Println("secondNumber2 (value)	:", secondNumber2)
	fmt.Println("secondNumber2 (memori address)	:", &secondNumber2)

	// Contoh 3
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value)	:", firstPerson)
	fmt.Println("firstPerson (memori address)	:", &firstPerson)
	fmt.Println("secondPerson (value)	:", *secondPerson)
	fmt.Println("secondPerson (memori address)	:", &secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Doe"

	fmt.Println("firstPerson (value)	:", firstPerson)
	fmt.Println("firstPerson (memori address)	:", &firstPerson)
	fmt.Println("secondPerson (value)	:", *secondPerson)
	fmt.Println("secondPerson (memori address)	:", &secondPerson)

	// Contoh 4
	var a int = 10

	fmt.Println("Before:", a)

	changeValue(&a)

	fmt.Println("After:", a)

	// 4. Struct
	// Contoh 1
	var employee Employee

	employee.name = "Airell"
	employee.age = 23
	employee.division = "Curriculum Developer"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)

	// Contoh 2
	var employee1 = Employee{}
	employee1.name = "Airell"
	employee1.age = 23
	employee1.division = "Curriculum Developer"

	var employee2 = Employee{name: "Ananda", age: 23, division: "Finance"}

	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)

	// Contoh 3
	var employee3 = Employee{name: "Airell", age: 23, division: "Curriculum Developer"}
	var employee4 *Employee = &employee3

	fmt.Println("Employee3 name:", employee3.name)
	fmt.Println("Employee4 name:", employee4.name)

	fmt.Println(strings.Repeat("#", 20))
	employee4.name = "Ananda"

	fmt.Println("Employee3 name:", employee3.name)
	fmt.Println("Employee4 name:", employee4.name)

	// Contoh 4
	var employee5 = EmployeeA{}

	employee5.person.name = "Airell"
	employee5.person.age = 23
	employee5.division = "Curriculum Developer"

	fmt.Printf("%+v", employee5)

	// Contoh 5
	// Anonymous struct tanpa pengisian field
	var employee6 = struct {
		person   Person
		division string
	}{}
	employee6.person = Person{name: "Airell", age: 23}
	employee6.division = "Curriculum developer"

	// Anonymous struct dengan pengisian field
	var employee7 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Ananda", age: 23},
		division: "Finance",
	}

	fmt.Printf("Employee6: %+v\n", employee6)
	fmt.Printf("Employee7: %+v\n", employee7)

	// Contoh 6
	var people = []Person{
		{name: "Airell", age: 23},
		{name: "Ananda", age: 23},
		{name: "Mailo", age: 23},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}

	// Contoh 7
	var employee8 = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Airell", age: 23}, division: "Curriculum Developer"},
		{person: Person{name: "Ananda", age: 23}, division: "Finance"},
		{person: Person{name: "Mailo", age: 21}, division: "Marketing"},
	}

	for _, v := range employee8 {
		fmt.Printf("%+v\n", v)
	}

	// 5. Method
	// Contoh 1
	var person2 = Person{name: "Airell", age: 23}
	fmt.Println(person2.Introduce("Hello everyone!!"))

	// Contoh 2
	var person3 = Person{name: "Airell", age: 23}

	person3.ChangeName1()
	fmt.Println("Change name with ChangeName1 method", person3.name)

	person3.ChangeName2()
	fmt.Println("Change name with ChangeName2 method", person3.name)

	// Contoh 3
	// Mengakses method pointer dari struct biasa
	var person4 = Person{name: "Airell", age: 23}
	person4.Greet()

	// Mengakses method pointer dari struct pointer
	var person5 = &Person{name: "Mailo", age: 15}
	person5.Greet()

	// 6. Reflect
	// Contoh 1
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}

	// Contoh 2
	var number2 = 23
	var reflectValue2 = reflect.ValueOf(number2)

	fmt.Println("tipe variabel :", reflectValue2.Type())
	fmt.Println("nilai variabel :", reflectValue2.Type())

	var nilai = reflectValue2.Interface().(int)
	fmt.Println("nilai :", nilai)

	// Contoh 3
	var s1 = &student{Name: "wick", Grade: 2}
	s1.getPropertyInfo()

	// Contoh 4
	var s2 = &student{Name: "john wick", Grade: 2}
	fmt.Println("nama :", s2.Name)

	var reflectValue3 = reflect.ValueOf(s2)
	var method = reflectValue3.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("nama :", s2.Name)

	// 7. Exported & Unexported
	// Contoh 1
	helpers.Greet()
	// Error
	// helpers.greet()

	// Contoh 2
	helpers.Greet()
	var person = helpers.PersonA{}
	person.Invokegreet()
}

// 1. Function
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
func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}

// Contoh 6
func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}

	return total
}

// Contoh 7
func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ",")

	fmt.Println("Hello there!!! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}

// Closure
// Contoh 1
func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s does'nt exist!!!", s)
		}
		return fmt.Sprintf("We found %s at position %d", s, position+1)
	}
}

// Contoh 2
func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}

// Contoh 3
func findOddNumbers2(numbers []int, callback isOddNum) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}

// Pointer
// Contoh 1
func changeValue(number *int) {
	*number = 20
}
