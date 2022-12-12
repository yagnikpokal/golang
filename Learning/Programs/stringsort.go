package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortstring(str string) string {
	s := strings.Split(str, "") //Convert string to slice
	// fmt.Println(s)
	sort.Strings(s) //Sort the slice
	// fmt.Println(s)
	//fmt.Println("This", strings.Join(s, ""))

	return strings.Join(s, "") //Convert slice to string

}
func main() {
	str := "yagnikpokal"
	sorted := sortstring(str)
	fmt.Println(str)
	fmt.Println(sorted)

}
