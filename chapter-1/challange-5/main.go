package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	data1 := []interface{}{"bisa1", "bisa2", "bisa3"}
	data2 := []interface{}{"coba1", "coba2", "coba3"}

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			for j := 0; j < 2; j++ {
				fmt.Println(data1, j+1)
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			for j := 0; j < 2; j++ {
				fmt.Println(data2, j+1)
			}
		}()
	}

	wg.Wait()
}
