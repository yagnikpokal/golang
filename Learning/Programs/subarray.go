/*
Given an array of integers, find the subarrays where the absolute difference between any two elements is less than or equal to 1.

Input Array =  [4,6,5,3,3,1]

Answer  =
[4,5]
[4,3,3]
*/package main

import (
	"fmt"
)

func findSubarrays(arr []int) [][]int {
	n := len(arr)
	result := [][]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if abs(arr[i]-arr[j]) <= 1 {
				result = append(result, []int{arr[i], arr[j]})
			}
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	arr := []int{4, 6, 5, 3, 3, 1}
	result := findSubarrays(arr)
	fmt.Println(result)
}
