/*
Write a function to read a slice of integer and
select the single number as output.
For examples :

Input: [2,2,1]
Output: 1

Input: [3,3,4,1,2,1,2]
Output: 4
*/

package main

import "fmt"

func findSingleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}
func main() {
	number1 := []int{2, 2, 1}
	single1 := findSingleNumber(number1)
	fmt.Println(single1)
	number2 := []int{3, 3, 4, 1, 2, 1, 2}
	single2 := findSingleNumber(number2)
	fmt.Println(single2)
}
