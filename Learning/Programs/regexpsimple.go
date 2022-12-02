package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "geeksforgeeks"
	value, err := regexp.MatchString("geeks", str)
	fmt.Println(value, err)
	value1, err := regexp.MatchString("yagnik", str)
	fmt.Println(value1, err)
}
