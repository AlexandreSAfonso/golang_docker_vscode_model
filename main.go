package main

import "fmt"


func sum(numbers []int) int {
	// Initialize a variable to hold the sum
	sum := 0
	// Iterate over the slice of integers
	for _, number := range numbers {
		// Add each number to the sum
		sum += number
	}
	// Return the total sum
	return sum
}

func main() {
	// Example usage of the function
	numbers := []int{1, 2, 3, 4, 5}
	result := sum(numbers)
	fmt.Println("The sum is:", result) // Output: The sum is: 15
}
