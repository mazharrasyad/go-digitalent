package main

import (
	"fmt"
	"math"
)

// 1. Interface
// Contoh 1
type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// Contoh 2
func print(t string, s shape) {
	fmt.Printf("%s area %v\n", t, s.area())
	fmt.Printf("%s perimeter %v\n", t, s.perimeter())
}

// Contoh 3
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {
	// 1. Interface
	// Contoh 1
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{width: 3, height: 2}

	fmt.Printf("Type of c1: %T", c1)
	fmt.Printf("Type of r1: %T", r1)

	// Contoh 2
	var c2 shape = circle{radius: 5}
	var r2 shape = rectangle{width: 3, height: 2}

	fmt.Println("Circle area", c2.area())
	fmt.Println("Circle perimeter", c2.perimeter())

	fmt.Println("Rectangle area", r2.area())
	fmt.Println("Rectangle perimeter", r2.perimeter())

	// Contoh 3
	var c3 shape = circle{radius: 5}
	var r3 shape = rectangle{width: 3, height: 2}

	print("Circle", c3)
	print("Rectange", r3)

	// Contoh 4
	var c4 shape = circle{radius: 5}

	// Error
	// c4.volume()
	c4.(circle).volume()

	// Contoh 5
	var c5 shape = circle{radius: 5}
	value, ok := c5.(circle)

	if ok == true {
		fmt.Printf("Circle value: %+v\n", value)
		fmt.Printf("Circle volume: %f", value.volume())
	}

	// 2. Empty Interface
	// Contoh 1
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan Sudirman"
	randomValues = 20
	randomValues = true
	randomValues = []string{"Airell", "Nanda"}

	// Contoh 2
	var v interface{}

	v = 20
	// Error
	// v = v * 9
	if value, ok := v.(int); ok == true {
		v = value * 9
	}

	// Contoh 3
	rs := []interface{}{1, "Airell", true, 2, "Ananda", true}
	rm := map[string]interface{}{
		"Name":   "Airell",
		"Status": true,
		"Age":    23,
	}

	_, _ = rs, rm
}
