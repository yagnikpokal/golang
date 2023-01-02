package main

import (
	"fmt"
	"sync"
)

type DbConnection struct{}

var (
	mut  sync.Mutex
	conn *DbConnection
)

func GetConnection() *DbConnection {
	// Lock and unlock the entire GetInstance function
	mut.Lock()
	defer mut.Unlock()

	if conn == nil {
		fmt.Println("Inside")
		conn = &DbConnection{}
	}
	fmt.Println("Outside")
	return conn
}
func main() {
	for i := 0; i < 5; i++ {
		_ = GetConnection()

	}
}

/*
go run oncewithmutex.go
Inside
Outside
Outside
Outside
Outside
Outside
*/
