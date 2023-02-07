package main

import (
	"fmt"
	"sync"
)

func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
	}
}

func main() {
	tasks := make(chan int)
	var wg sync.WaitGroup

	// Create a fixed number of worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Send tasks to the task channel
	for i := 1; i <= 5; i++ {
		tasks <- i
	}
	close(tasks)

	// Wait for all the worker goroutines to finish
	wg.Wait()
	fmt.Println("All tasks have been processed")
}

/*
 go run workerpoolwaitgroup3worker5tasksinglechannel.go
Worker 3 processing task 1
Worker 3 processing task 4
Worker 3 processing task 5
Worker 2 processing task 3
Worker 1 processing task 2
All tasks have been processed

*/
