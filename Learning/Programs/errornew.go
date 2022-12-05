package main

import (
	"errors"
	"fmt"
)

func mul(i int, j int) (int, error) {
	if i == 0 || j == 0 {
		return 0, errors.New("0 can not be multipled")
	}
	return i * j, nil
}

func main() {
	d, err := mul(2, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}

}
