package main

import "fmt"

func main() {
	char := "a"
	str := "yagnikpokal"
	counter := 0
	for i := 0; i <= len(str)-1; i++ {
		if string(str[i]) == char {
			counter++
		}
	}
	fmt.Println(counter)
}
