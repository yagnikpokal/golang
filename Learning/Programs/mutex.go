package main

import (
	"fmt"
	"sync"
)

var counter = 0

func increment(m *sync.Mutex, wg *sync.WaitGroup) {
	m.Lock()
	counter = counter + 1
	m.Unlock()
	wg.Done()
}

func main() {
	var m sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&m, &wg)
	}
	wg.Wait()
	fmt.Println(counter)

}
