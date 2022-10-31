package main

import "fmt"

func main() {

	fmt.Println("Begning")

	defer fmt.Println("One")

	defer fmt.Println("Two")

	defer fmt.Println("Three")

	fmt.Println("End")

}
