// Take array
// 10000, 100000
// if it is 5 digit

package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{540, 12345678, 9}

	for _, num := range arr {
		if len(strconv.Itoa(num)) <= 5 {

			sum := 0
			for num > 0 {
				sum += num % 10
				num /= 10
			}
			for i := num; i > 0; i = i / 10 {
				sum += i % 10
			}
			fmt.Println(sum)
		} else {
			product := 1
			for i := num; i > 0; i = i / 10 {
				product *= i % 10
			}
			fmt.Println(product)
		}
	}
}

/*
func sumOfnum(n int)int{
	sum:=0
	num:=strconv.Itoa(n)
	for _,digit:=range num{
		d,_:=strconv.Atoi
	}
}
func main(){
	number:=[]int{12345,678,9014}
	//var sum int
	var mul int
	for _,n:=range number {
		if len(number[n]<=5){
//Do sum of number
			sum := 0
			for i := v; i > 0; i = i / 10 {
				sum += i % 10
			}
			fmt.Printf("The sum is %d", v, sum)
fmt.Println(sum)
		}else{
			product :=1
			for i:=
			//multiplication
					product := 1
			for i := v; i > 0; i = i / 10 {
				product *= i % 10
			}
			fmt.Printf("The multiplication is %d", v, product)
		}
	}
}*/
