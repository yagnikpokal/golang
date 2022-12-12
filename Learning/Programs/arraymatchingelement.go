package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 5, 3, 7, 5, 4, 7, 7, 7}
	count := 1
	for i := 0; i < len(arr)-1; i++ {

		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[i] == arr[j] {
				fmt.Println(arr[j])
				count++
			}
		}
	}
}
