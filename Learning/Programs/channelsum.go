// Go program to illustrate send
// and receive operation
package main

import "fmt"

func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}
func main() {

	// Creating a channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23

}

/*package main

import "fmt"

var number = 0

func increment(mychan chan int) {
	fmt.Println(50 + <-mychan)
}
func main() {
	mychan := make(chan int)
	go increment(mychan)
	mychan <- 50

}
*/
