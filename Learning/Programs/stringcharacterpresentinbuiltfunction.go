package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "yagnikpokal"
	present := strings.Contains(str, "y") // Single character present or not
	fmt.Println(present)
	wordpresent := strings.Contains(str, "yag") // Word present or not
	fmt.Println(wordpresent)
}
