package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "yagnikpokal123#/*"
	for _, i := range str {
		if unicode.IsDigit(i) {
			fmt.Println(string(i), "is Digit")
		} else if unicode.IsLetter(i) {
			fmt.Println(string(i), "s Letter")
		} else {
			fmt.Println(string(i), "special character")
		}
	}
}
