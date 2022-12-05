package main

import (
	"errors"
	"fmt"
)

func returnError() (int, error) { // declare return type here
	return 42, errors.New("Error occured!") // return it here
}

func main() {
	v, e := returnError()
	if e != nil {
		fmt.Println(e, v) // Error occured! 42
	}
}
