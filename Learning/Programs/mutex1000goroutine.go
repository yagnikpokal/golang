package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter uint64
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

/*
go run mutex1000goroutine.go
1000

go run -race mutex1000goroutine.go
1000
*/
