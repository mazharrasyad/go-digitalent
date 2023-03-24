package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 1. Goroutines
	// Contoh 1
	go goroutine()

	// Contoh 2
	fmt.Println("main execution started")

	go firstProcess(8)
	secondProcess(8)

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	fmt.Println("main execution ended")

	// Contoh 3
	fmt.Println("main execution started")

	go firstProcess(8)
	secondProcess(8)

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	time.Sleep(time.Second * 2)
	fmt.Println("main execution ended")

	// 2. Sync - Waitgroup
	fruits := []string{"apple", "manggo", "durian", "rambutan"}

	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1)
		go printFruit(index, fruit, &wg)
	}

	wg.Wait()

	// 3. Channels
	// Contoh 1
	c := make(chan string)
	// Mengirim data kepada channel
	// c <- value
	// Menerima data dari channel
	// result := <-c

	go introduce("Airell", c)
	go introduce("Nanda", c)
	go introduce("Mailo", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	close(c)

	// Contoh 2
	c2 := make(chan string)
	students := []string{"Airell", "Mailo", "Indah"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("Student", student)
			result := fmt.Sprintf("Hai, my name is %s", student)
			c2 <- result
		}(v)
	}

	for i := 1; i <= 3; i++ {
		print(c2)
	}

	close(c2)

	// Contoh 3
	c3 := make(chan string)

	students2 := []string{"Airell", "Mailo", "Indah"}

	for _, v := range students2 {
		go introduce2(v, c3)
	}

	for i := 1; i <= 3; i++ {
		print(c3)
	}

	close(c3)

	// Contoh 4
	// Unbuffered Channel
	c4 := make(chan int)

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c4)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	d := <-c4
	fmt.Println("main goroutine received data:", d)

	close(c4)
	time.Sleep(time.Second)

	// Contoh 5
	c5 := make(chan int, 3) // Buffered Channel

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Println("func goroutine starts sending data into the channel")
			c <- 10
			fmt.Println("func goroutine after sending data into the channel")
		}

		close(c)
	}(c5)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c5 {
		fmt.Println("main goroutine received value from channel:", v)
	}

	// Contoh 6
	c6 := make(chan string)
	c7 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		c6 <- "Hello!"
	}()

	go func() {
		time.Sleep(1 * time.Second)

		c7 <- "Salut!"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c6:
			fmt.Println("Received", msg1)
		case msg2 := <-c7:
			fmt.Println("Received", msg2)
		}
	}
}

// 1. Goroutines
// Contoh 1
func goroutine() {
	fmt.Println("Hello")
}

// Contoh 2
func firstProcess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First process func ended")
}

func secondProcess(index int) {
	fmt.Println("Second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("Second process func ended")
}

// 2. Sync - Waitgroup
func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index => %d, fruit => %s\n", index, fruit)
	wg.Done()
}

// 3. Channel
// Contoh 1
func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hai, my name is %s", student)
	c <- result
}

// Contoh 2
func print(c <-chan string) {
	fmt.Println(<-c)
}

func introduce2(student string, c chan<- string) {
	result := fmt.Sprintf("Hai, my name is %s", student)
	c <- result
}
