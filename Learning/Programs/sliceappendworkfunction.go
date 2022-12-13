package main

import "fmt"

func ModifyData(a []int) {
	a[0] = 5
}

func AddData(a []int) []int {
	a = append(a, 4)
	return a
}

func main() {

	a := []int{1, 2, 3} // Slice
	a = AddData(a)      //{1,2,3,4}
	fmt.Println(a)      // 1,2,3,4
	ModifyData(a)       //5,2,3,4
	fmt.Println(a)      // 5,2,3,4

}
