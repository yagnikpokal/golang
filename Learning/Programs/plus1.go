package main

import "fmt"

/*func plusone(arr []int) {

}*/
func main() {
	arr := []int{1, 4, 1, 3}
	fmt.Println(arr)
	D := len(arr) - 1

	for i := D; i >= 0; i-- {
		if arr[D] == 9 {
			arr[i] = 0
			D--

		} else {
			arr[i] = arr[i] + 1
			break
			D--
		}

	}

	fmt.Println(arr)

}
