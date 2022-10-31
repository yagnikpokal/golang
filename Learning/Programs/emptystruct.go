package main

import (
	"fmt"
	"unsafe"
)

func main() {

	a := struct{}{}
	fmt.Println(unsafe.Sizeof(a))
}
