// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

var a = 0

func increment(wg *sync.WaitGroup) {
	a = a + 1
	wg.Done()
}

func main() {
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}

	w.Wait()

	fmt.Println(a)
}
