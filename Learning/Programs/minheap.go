package main

import "fmt"

func heapify(heap *[]int, i int) {
	smallest := i
	lChild := 2*i + 1
	rChild := 2*i + 2

	if lChild < len(*heap) && (*heap)[lChild] < (*heap)[smallest] {
		smallest = lChild
	}
	if rChild < len(*heap) && (*heap)[rChild] < (*heap)[smallest] {
		smallest = rChild
	}

	if smallest != i {
		(*heap)[i], (*heap)[smallest] = (*heap)[smallest], (*heap)[i]
		heapify(heap, smallest)
	}
}
func main() {
	input := []int{4, 12, 3, 6, 5}
	fmt.Println(input)
	heapify(&input, 4)
	fmt.Println(input)
}
