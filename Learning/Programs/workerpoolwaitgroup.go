package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		//time.Sleep(time.Second)  this will run a job sequnceially since we use the 1 second time to cmplete the go routine
		results <- j * 2
	}
}

func main() {
	job := make(chan int, 10)
	result := make(chan int, 10)
	var wg sync.WaitGroup

	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go worker(w, job, result, &wg)
	}
	for j := 1; j <= 9; j++ {
		job <- j
	}
	close(job)
	for a := 1; a <= 9; a++ {
		<-result
	}

	wg.Wait()
	fmt.Println("All worker have been processed")
}

/*
go run workerpoolwaitgroup.go
worker 2 processing job 1
worker 2 processing job 3
worker 2 processing job 4
worker 2 processing job 5
worker 2 processing job 6
worker 2 processing job 7
worker 2 processing job 8
worker 2 processing job 9
worker 1 processing job 2
All worker have been processed

*/
