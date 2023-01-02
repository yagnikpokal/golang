package main

import "fmt"

func linearsearch(arr []int, digit int) int {
	for index, value := range arr {
		if value == digit {
			return index
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 5, 7, 8}
	digit := 2
	index := linearsearch(arr, digit)
	fmt.Println(index)

}
