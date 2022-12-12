package main

import "fmt"

func main() {
	str := "yagnikpokal"
	counter := 0
if 'a'<=

	for i := 0; i < len(str)-1; i++ {
		if string(str[i]) == "a" || string(str[i]) == "e" || string(str[i]) == "i" || string(str[i]) == "o" || string(str[i]) == "u" {
			counter++
		}
	}


	fmt.Println("vowels are", counter)
	fmt.Println("consonents are", len(str)-counter)

}
