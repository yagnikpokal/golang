package main

import "fmt"

func main() {
	original_string := "madam1"
	var reverse_string = ""
	for i := len(original_string) - 1; i >= 0; i-- {
		reverse_string += string(original_string[i])

	}
	if original_string == reverse_string {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}

}
