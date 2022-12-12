package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortstring(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func main() {
	str1 := "yagnikpokal"
	str2 := "pokalyagnik"

	sorted1 := sortstring(str1)
	sorted2 := sortstring(str2)
	if sorted1 == sorted2 {
		fmt.Println("string is anagram")
	} else {
		fmt.Println("string is not anagram")

	}

}
