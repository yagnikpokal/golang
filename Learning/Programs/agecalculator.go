package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	myBirth := "1994-mar-05"
	layout := "2006-Jan-02"
	myBirthdate, _ := time.Parse(layout, myBirth)
	age := currentTime.Sub(myBirthdate).Hours() / 8760 // 8760 is convert hour to year
	fmt.Println(age)
}
