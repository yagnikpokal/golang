package main

import (
	"fmt"
)

//Simple example for this in terms of push/pop

var stack []int

func Stack() {

	stack = append(stack, 10)
	stack = append(stack, 20)
	stack = append(stack, 30)

	for len(stack) > 0 {

		n := len(stack) - 1
		stack = stack[:n]
		fmt.Println(stack)

	}
}

func main() {
	Stack()

}
