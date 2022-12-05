package main

import (
	"errors"
	"fmt"
)

func mul(i int, j int) (int, error) {
	if i == 0 || j == 0 {
		return 42, errors.New("0 can not be multipled")
	}
	return i * j, nil
}

func main() {
	d, _ := mul(2, 0)
	fmt.Println(d)
}
