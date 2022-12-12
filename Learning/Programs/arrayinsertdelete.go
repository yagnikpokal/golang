package main

import "fmt"

func main() {
	arr := []int{10, 52, 32, 46}
	fmt.Println(arr)
	arr[1] = 16 // Copying/inserting the value
	fmt.Println(arr)
	arr = append(arr, 0) // Making space for the new element
	fmt.Println(arr)

	copy(arr[3:], arr[2:]) // copy + Shifting elements
	fmt.Println(arr)

	arr = append(arr[:2], arr[3:]...) //This will delete second element
	fmt.Println(arr)

	index := 3
	arr = append(arr[:index], arr[index+1:]...) //This will delete Third element
	fmt.Println(arr)

}
