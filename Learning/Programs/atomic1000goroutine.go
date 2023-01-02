package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddUint64(&counter, 1)
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println(counter)
}

/*
go run atomic1000goroutine.go
1000
*/
