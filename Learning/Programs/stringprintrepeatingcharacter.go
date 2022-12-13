// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "yagnikpokal"
	var mymap = make(map[string]int)
	for I, j := range name {

		fmt.Println(I, string(j))
		singlecharacter := string(j)
		sumcharacter := strings.Count(name, singlecharacter)
		//fmt.Println(string(j), "repeated", sumcharacter, "times")
		mymap[string(j)] = sumcharacter
	}
	fmt.Println(mymap) // To print all the values

	for key, element := range mymap {
		if element >= 2 {
			fmt.Println(key, "Present", element, "times") // To print only repeating character
		}
	}

}
