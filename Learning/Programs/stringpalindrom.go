package main

import (
	"fmt"
)

func main() {
	str := "nitin"
	reverse := ""

	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	fmt.Println(str)
	fmt.Println(reverse)
	if str == reverse {
		fmt.Println("Palindrome")
	}

}
