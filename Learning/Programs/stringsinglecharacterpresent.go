package main

import "fmt"

func main() {
	myname := "yagnikpokal"
	character := "g"

	for _, j := range myname {
		if string(j) == character {
			fmt.Println("g is present in the string")
		}
	}

}
