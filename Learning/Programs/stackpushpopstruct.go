package main

import "fmt"

// Stack can be accessed or in data or added data by using struct with slice
type Stack struct {
	items []int
}

// Create the 2 methods push and pop
// Push will add a value at the end
func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

//POP will remove value at the end
// and return the removed
func (s *Stack) pop() {
	s.items = s.items[:len(s.items)-1]
}

func main() {

	// Creating the stack
	myStack := Stack{}
	fmt.Println(myStack)

	// Push the items add the items in stack
	myStack.push(100)
	myStack.push(200)
	myStack.push(300)
	fmt.Println(myStack)

	// Pop the items remove the items from the stack
	myStack.pop()
	myStack.pop()

	fmt.Println(myStack)

}
