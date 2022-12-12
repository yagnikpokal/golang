package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 6, 5, 4}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}
