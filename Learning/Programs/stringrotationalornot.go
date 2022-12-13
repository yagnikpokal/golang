package main

import (
	"fmt"
	"strings"
)

func isSubstring(s1, s2 string) bool {
	if !strings.Contains(s1, s2) {
		return false
	}
	return true
}

func rotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	} else {
		s1 = s1 + s1
		return isSubstring(s1, s2)
	}
	return true
}

func main() {

	fmt.Println(rotation("yagnik", "agniky"))
}
