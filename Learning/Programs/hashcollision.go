package main

import (
	"fmt"
)

func hash(s string) int {
	// This is a very simple hashing algorithm that just adds the ASCII values
	// of the characters in the string
	var hash int
	for i := 0; i < len(s); i++ {
		hash += int(s[i])
	}
	return hash
}

func main() {
	// Find two different strings that have the same hash value
	var s1, s2 string
	for {
		s1 = "hello"
		s2 = "world"
		if hash(s1) == hash(s2) {
			break
		}
	}

	// Print the strings and their hash values
	fmt.Println("s1:", s1, "hash(s1):", hash(s1))
	fmt.Println("s2:", s2, "hash(s2):", hash(s2))
}
