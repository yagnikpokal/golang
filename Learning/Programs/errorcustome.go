package main

import (
	"fmt"
	"os"
)

type MyError struct{}

func (m *MyError) Error() string {
	return "boom"
}

func sayHello() (string, error) {
	return "", &MyError{}
}
func main() {
	_, err := sayHello()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
