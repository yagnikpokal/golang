package main

import (
	"fmt"
	"sync"
)

var counter = 0

func main() {
	var wg sync.WaitGroup
	mychan := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mychan <- true
			counter++
			<-mychan
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

/*
go run -race channel1000goroutine.go
1000
*/
