package main

import "fmt"

func isAnagram(s string, t string) bool {

	lenS := len(s)
	lenT := len(t)

	if lenS != lenT {
		return false
	}

	anagramMap := make(map[string]int)

	for i := 0; i < lenS; i++ {
		anagramMap[string(s[i])]++
	}

	for i := 0; i < lenT; i++ {
		anagramMap[string(t[i])]--
	}

	for i := 0; i < lenS; i++ {
		if anagramMap[string(s[i])] != 0 {
			return false
		}
	}

	return true
}

func main() {
	output := isAnagram("abc", "bac")
	fmt.Println(output)

	output = isAnagram("abc", "bc")
	fmt.Println(output)
}
