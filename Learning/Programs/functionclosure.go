package main

import "fmt"

func main() {
	counter := 0
	increment := func() int {
		counter++
		return counter
	}
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}

/*
A closure in Go is a function that refers to variables from its surrounding outside of the scope.
When a closure is created, it captures the values of these variables at the time of creation and retains them for later use, even if the original variables go out of scope.
1
2
*/
