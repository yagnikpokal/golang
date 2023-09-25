package main

import "fmt"

func isPalindrom(num int) bool {
	original := num
	reverse := 0
	for num > 0 {
		reminder := num % 10
		reverse = reverse*10 + reminder
		num /= 10
	}
	return original == reverse
}
func main() {

	if isPalindrom(123) {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Not palindrom")
	}
}
