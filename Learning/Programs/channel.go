package main

import (
	"fmt"
	"sync"
)

var counter = 0

func increment(wg *sync.WaitGroup, mychan chan bool) {
	mychan <- true
	counter = counter + 1
	<-mychan
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	mychan := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, mychan)
	}
	wg.Wait()
	fmt.Println(counter)
}
