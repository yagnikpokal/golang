package main

import (
	"fmt"
	"sync"
)

type DbConnection struct{}

var (
	dbConnOnce sync.Once
	conn       *DbConnection
)

func GetConnection() *DbConnection {
	dbConnOnce.Do(func() {
		conn = &DbConnection{}
		fmt.Println("Inside")
	})
	fmt.Println("Outside")
	return conn
}

func main() {
	for i := 0; i < 5; i++ {
		_ = GetConnection()
	}
}

/*
go run once.go
Inside
Outside
Outside
Outside
Outside
Outside

*/
