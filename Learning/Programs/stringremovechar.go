package main

import "fmt"

func main() {
	str := "yagnikpokal"
	char := "a"
	str2 := ""

	for i := 0; i < len(str)-1; i++ {
		if string(str[i]) == string(char) {

		} else {
			str2 += string(str[i])
		}
	}
	fmt.Println(str2)
}
