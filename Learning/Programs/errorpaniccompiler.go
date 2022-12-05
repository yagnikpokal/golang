package main

import "fmt"

func foo() {

	defer panic("There is something wrong!")
	fmt.Println("hello from the deferred function!")
}
func main() {
	foo()
}
