package main

import (
	"fmt"
	"sync"
)

func addnumbers(wg *sync.WaitGroup, number chan int) {
	for i := 1; i <= 10; i++ {
		number <- i
	}
	close(number)
	wg.Done()
}

func readnumber(wg *sync.WaitGroup, number chan int) {
	for n := range number {
		fmt.Println(n)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	number := make(chan int)
	wg.Add(2)
	go addnumbers(&wg, number)
	go readnumber(&wg, number)
	wg.Wait()
}

//   function1 add number
//  Calculate sum(1 … 10) in one routine and Print the sum in another routine

//
/*
SELECT MAX(marks) FROM students_marks WHERE marks NOT IN(SELECT max(marks)) FROM students_marks
WITH CTE as(

)

*/
