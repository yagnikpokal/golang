package main

import "fmt"

func main() {
	arr := []int{3, 3}
	num := 3
	fmt.Println(arr)
	j := 0
	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] == num {
			j++
		}
	}
	fmt.Println(j)
	for k := 0; k < j; k++ {
		for i := 0; i <= len(arr)-1; i++ {
			if arr[i] == num {
				arr = append(arr[:i], arr[i+1:]...)
			}
		}
	}

	fmt.Println(arr)

}
