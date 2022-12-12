package main // This is not a perfect example need to check more on this

import (
	"fmt"
	"time"
)

func server1(channel1 chan string) {
	for {
		time.Sleep(3 * time.Second)
		channel1 <- "Echo from server1"
	}
}
func server2(channel2 chan string) {
	for {
		time.Sleep(3 * time.Second)
		channel2 <- "Echo from server2"
	}
}

// Main function
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go server1(channel1)
	go server2(channel2)

	for {
		select {
		case o := <-channel1:
			fmt.Println(o)
		case t := <-channel2:
			fmt.Println(t)
		default:
			// fmt.Println("No data to receive")
			// time.Sleep(3 * time.Second)
		}
	}
}
