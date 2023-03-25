package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex

func main() {
	data1 := []string{"bisa1", "bisa2", "bisa3"}
	data2 := []string{"coba1", "coba2", "coba3"}

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("===== Secara Acak =====")
	go printRandom(data1, &wg)
	go printRandom(data2, &wg)

	wg.Wait()

	fmt.Println("\n===== Secara Rapih =====")
	wg.Add(2)

	go printMutex(data1, &wg)
	go printMutex(data2, &wg)

	wg.Wait()
}

func printRandom(data interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 4; i++ {
		fmt.Printf("%s %d\n", data, rand.Intn(3)+1)
		time.Sleep(time.Millisecond * 500)
	}
}

func printMutex(data interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 4; i++ {
		mutex.Lock()
		fmt.Printf("%s %d\n", data, i)
		mutex.Unlock()
		time.Sleep(time.Millisecond * 500)
	}
}
