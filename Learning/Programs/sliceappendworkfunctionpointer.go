package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")

	}
}

func main() {

	a := []int{8, 6, 7, 5, 3, 0, 9, -1, 3, 4}
	// If i use a := []int{8, 6, 7, 5, 3, 0, 9, -1} then it works fine

	c := make(chan int)

	go printCount(c)

	for _, v := range a {

		c <- v

	}

	time.Sleep(time.Second * 1)
	fmt.Println("End of main")
}

/*
package main

import "fmt"

func ModifyData(a []int) {
	a[0] = 5
}

func AddData(a *[]int) {
	*a = append(*a, 4)
}

func main() {

	a := []int{1, 2, 3} // Slice
	AddData(&a)         //{1,2,3,4}
	fmt.Println(a)      // 1,2,3,4
	ModifyData(a)       //5,2,3,4
	fmt.Println(a)      // 5,2,3,4

}*/
