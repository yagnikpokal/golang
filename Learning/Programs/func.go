// string as input
// check string is palindro or not
package main

import (
	"fmt"
	"strings"
)

func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {

	s = strings.ToLower(s)

	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}
		for i < j && !isChar(s[j]) {
			j--

		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	Y := "yagnik"
	D := isPalindrome(Y)
	fmt.Println(D)
}
