package main

// findChampion function finds the team that will be the champion of the tournament.
func findChampion(grid [][]int) int {
	// Step 1: Initialize `champion` to 0.
	// Assume that the first team (team 0) is the strongest for now.
	champion := 0

	// Step 2: Iterate over all the teams.
	// Check if any other team can beat the current `champion`.
	for i := 1; i < len(grid); i++ {
		// If team `i` beats the current `champion`, update `champion` to `i`.
		// This is because team `i` is stronger than the current `champion`.
		if grid[i][champion] == 1 {
			champion = i
		}
	}

	// Step 3: Return the final `champion`.
	return champion
}
