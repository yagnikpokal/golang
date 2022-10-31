package main

import "fmt"

func sum(value ...int) int {
	sum := 0
	for _, j := range value {
		sum += j
	}
	return sum
}
func main() {
	x := []int{0, 1, 5, 6, 8, 6}
	y := sum(x...)
	fmt.Println(y)
}
