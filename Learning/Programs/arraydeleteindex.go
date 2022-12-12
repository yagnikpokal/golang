package main

import "fmt"

func delete_index(arr []int, a int) []int {
	return append(arr[:a], arr[a+1:]...)
	//return arr
}

func main() {
	arr := []int{10, 5, 6, 7, 9, 2}
	fmt.Println(arr)
	arr = delete_index(arr, 2) // Delete the index no 2
	fmt.Println(arr)
}
