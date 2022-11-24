package main

import "fmt"

// Queue can be accessed or in data or added data by using struct with slice
type Queue struct {
	items []int
}

// Create the 2 methods Enqueue and Dequeue
// Enqueue will add a value at the end

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

//Dequeue will remove value at the front
// and return the removed
func (q *Queue) Dequeue() {
	q.items = q.items[1:]
}
func main() {
	// Creating the queue
	myQueue := Queue{}
	fmt.Println(myQueue)

	// Enqueue add the items in queue
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)

	// Dequeue rremove the items from the queue
	myQueue.Dequeue()
	fmt.Println(myQueue)

}
