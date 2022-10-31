package main

import (
	"errors"
	"fmt"
)

var negativenumber = errors.New("Error : The number is negative")

func Sum(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, negativenumber
	}
	return a + b, nil
}

func main() {
	D, error := Sum(5, -65)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(D)
	}

}
