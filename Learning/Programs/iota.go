package main

import "fmt"

const (
	north = iota
	south
	east
	west
)

func main() {
	fmt.Println(north, south, east, west)

}
