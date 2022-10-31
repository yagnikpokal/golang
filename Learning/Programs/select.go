package main

import (
	"fmt"
	"time"
)

// Main function
func main() {
	one := make(chan int)
	two := make(chan int)

	for {
		select {
		case o := <-one:
			fmt.Println("One", o)

		case t := <-two:
			fmt.Println("two", t)

		default:
			fmt.Println("No data to receive")
			time.Sleep(50 * time.Millisecond)
		}
	}

}
