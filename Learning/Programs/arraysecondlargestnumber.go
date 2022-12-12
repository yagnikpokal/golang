package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{0, 2, 3, 5, 4, 15, 13, 6, 7}
	D := arr
	sort.Ints(D)
	fmt.Println(D[len(D)-2])
}
