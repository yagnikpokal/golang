package main

import (
	"fmt"
	"time"
)

func listentochannel(ch chan int) {
	for {

		i := <-ch
		fmt.Println(i)
		time.Sleep(1 * time.Second)

	}
}
func main() {
	ch := make(chan int)

	go listentochannel(ch)

	for i := 0; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
