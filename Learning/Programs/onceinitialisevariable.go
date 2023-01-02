package main

import (
	"fmt"
	"sync"
)

var globalVar int
var once sync.Once

func initialize() {
	globalVar = 42
}

func main() {
	go once.Do(initialize)
	go once.Do(initialize)
	once.Do(initialize)

	fmt.Println(globalVar) // prints 42
}

/*
Here we initialise variable 2 times but it will do only one time
go run onceinitialisevariable.go
42

*/
