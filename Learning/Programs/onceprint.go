package main

import (
	"fmt"
	"sync"
)

func setup() {
	fmt.Println("Initializing shared resource...")
}

func main() {
	var once sync.Once
	once.Do(setup)
	once.Do(setup)
	once.Do(setup)
}

/*
go run onceprint.go
Initializing shared resource...

*/
