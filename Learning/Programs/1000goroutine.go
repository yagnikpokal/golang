package main

import (
	"fmt"
	"sync"
)

var counter = 0

func increment(wg *sync.WaitGroup) {
	counter = counter + 1
	wg.Done()
}
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println(counter)
}
