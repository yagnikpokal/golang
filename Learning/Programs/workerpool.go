package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		//time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	job := make(chan int, 10)
	result := make(chan int, 10)
	for w := 1; w <= 2; w++ {
		go worker(w, job, result)
	}
	for j := 1; j <= 9; j++ {
		job <- j
	}
	close(job)
	for a := 1; a <= 9; a++ {
		<-result
	}
}

/*
go run workerpool.go
worker 2 processing job 1
worker 1 processing job 2
worker 1 processing job 3
worker 2 processing job 4
worker 2 processing job 5
worker 1 processing job 6
worker 1 processing job 7
worker 2 processing job 8
worker 2 processing job 9

*/
