package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "yagnikpokal"
	charcount := strings.Count(str, "a") // Count the character
	fmt.Println(charcount)
	wordcount := strings.Count(str, "yag") // Count the word
	fmt.Println(wordcount)

}
