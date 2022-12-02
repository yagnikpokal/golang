package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	re, _ := regexp.Compile(" ")
	replace := re.ReplaceAllString("my name is yagnik", "+")
	fmt.Println(replace)

	// Replace all the characters to uppercase using the function
	re1, _ := regexp.Compile("[aeiou]+")
	replace1 := re1.ReplaceAllStringFunc("My name is yagnik", strings.ToUpper)
	fmt.Println(replace1)
}
