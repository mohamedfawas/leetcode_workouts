package main

// findErrorNums finds the duplicate and missing numbers in a slice
func findErrorNums(nums []int) []int {
	// Get the length of the input slice
	n := len(nums)

	// Create a boolean slice to keep track of seen numbers
	// The size is n+1 because the numbers are from 1 to n
	seen := make([]bool, n+1)

	// Initialize variables
	var sum int       // To keep track of the sum of all numbers in nums
	var duplicate int // To store the number that appears twice
	var missing int   // To store the number that's missing

	// Calculate the expected sum of numbers from 1 to n
	expectedSum := n * (n + 1) / 2

	// Iterate through each number in the input slice
	for _, num := range nums {
		// If we've seen this number before, it's the duplicate
		if seen[num] {
			duplicate = num
		}

		// Mark this number as seen
		seen[num] = true

		// Add this number to our running sum
		sum += num
	}

	// Calculate the missing number
	// The sum we calculated includes the duplicate and excludes the missing number
	// So, (sum - duplicate) gives us the sum of all numbers except the missing one
	// Subtracting this from the expected sum gives us the missing number
	missing = expectedSum - (sum - duplicate)

	// Return the duplicate and missing numbers as a slice
	return []int{duplicate, missing}
}
