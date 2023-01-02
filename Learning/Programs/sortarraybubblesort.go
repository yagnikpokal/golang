package main

import "fmt"

func bubblesort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
func main() {
	arr := []int{4, 5, 6, 98, 7, 4, 3}
	sorted := bubblesort(arr)
	fmt.Println(arr)
	fmt.Println(sorted)

}
