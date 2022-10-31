// Go program to illustrate how to
// find the length of the channel

package main

import "fmt"

// Main function
func main() {

	// Creating a channel
	// Using make() function
	mychnl := make(chan string, 6)
	mychnl <- "GFG"
	mychnl <- "gfg"
	mychnl <- "Geeks"
	mychnl <- "GeeksforGeeks"

	// Finding the length of the channel
	// Using len() function
	fmt.Println("Capacity of the channel is: ", cap(mychnl))
}
