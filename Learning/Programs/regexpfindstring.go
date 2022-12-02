package main

import (
	"fmt"
	"regexp"
)

func main() {
	re2, _ := regexp.Compile("[0-9]+-y.*g") // This will print the string when first char is y and got second char g
	extract1 := re2.FindString("1994-yagnik_pokal")
	fmt.Println(extract1)
}
