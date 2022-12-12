package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{5, 4, 7, 9, 5, 3, 54, 6, 5, 69, 2}
	fmt.Println(arr)
	sort.Ints(arr)
	fmt.Println(arr)

	stringarr := []string{"c", "y", "a"}
	fmt.Println(stringarr)
	sort.Strings(stringarr)
	fmt.Println(stringarr)

}
