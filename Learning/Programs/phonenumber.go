package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	letters := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	result := []string{""}
	for i := range digits {
		digit := digits[i] - '0'
		if digit < 2 || digit > 9 {
			continue
		}
		letter := letters[digit]
		tmp := []string{}
		for j := range result {
			for k := range letter {
				tmp = append(tmp, result[j]+string(letter[k]))
			}
		}
		result = tmp
	}
	return result
}
func main() {
	myslice := letterCombinations("23")
	fmt.Println(myslice)
}

/*
Example 1:

Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
Example 2:

Input: digits = ""
Output: []
Example 3:

Input: digits = "2"
Output: ["a","b","c"]

*/
