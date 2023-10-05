package main

import "fmt"

func isBalanced(arr []int) bool {
	// Calculate the total sum of the array
	totalSum := 0
	for _, num := range arr {
		totalSum += num
	}

	// Check if the total sum is even; if it's odd, the array cannot be balanced
	if totalSum%2 != 0 {
		return false
	}

	// Calculate half of the total sum
	halfSum := totalSum / 2

	// Initialize a variable to keep track of the running sum of elements in the first half
	currentSum := 0

	// Iterate through the array, adding elements to currentSum until it exceeds or reaches halfSum
	for _, num := range arr {
		currentSum += num
		if currentSum == halfSum {
			return true
		}
	}

	// If we reach this point, the array is not balanced
	return false
}

func main() {
	// Example arrays
	balancedArray := []int{1, 2, 2, 1}
	unbalancedArray := []int{1, 2, 3, 4, 5}

	// Check if the arrays are balanced
	fmt.Println("Is balanced:", isBalanced(balancedArray))   // Should print true
	fmt.Println("Is balanced:", isBalanced(unbalancedArray)) // Should print false
}
