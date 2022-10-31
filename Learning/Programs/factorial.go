package main

import "fmt"

func factorial(a int) int {
	if a == 1 || a == 0 {
		return a
	}
	return a * factorial(a-1)
}

func main() {
	D := factorial(5)
	fmt.Println(D)
}
