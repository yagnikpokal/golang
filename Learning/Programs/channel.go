// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func sendfromchannel(ch chan int) {
	ch <- 47
}

func main() {
	ch := make(chan int)
	go sendfromchannel(ch)
	fmt.Println(<-ch)
}

/*
go run channel.go
42
*/
