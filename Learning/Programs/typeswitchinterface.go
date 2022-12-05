// Go program to illustrate type switch
package main

import "fmt"

func myfun(a interface{}) {

	// Using type switch
	switch a.(type) {

	case int:
		fmt.Println("Type: int, Value:", a.(int))
	case string:
		fmt.Println("\nType: string, Value: ", a.(string))
	case float64:
		fmt.Println("\nType: float64, Value: ", a.(float64))
	default:
		fmt.Println("\nType not found")
	}
}

// Main method
func main() {

	myfun("GeeksforGeeks")
	myfun(67.9)
	myfun(true)
}
