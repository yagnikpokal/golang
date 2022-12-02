package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, _ := regexp.Compile("geeks")
	str := "geeksforgeeks"
	myIndex := re.FindStringIndex(str) //Shows first index of the charcter matching string
	fmt.Println(myIndex)
	myIndex1 := re.FindAllStringSubmatchIndex("geeks for geeks", -1) // Shows the first and last character index of the matching string
	fmt.Println(myIndex1)

}
