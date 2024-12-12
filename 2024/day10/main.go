package main

import (
	"adventofcode-2024/models"
	"adventofcode-2024/utils"
)

// For all practical purposes
// this means that a hiking trail is any path that starts at height 0, ends at height 9
// and always increases by a height of exactly 1 at each step.
// Hiking trails never include diagonal steps
// only up, down, left, or right (from the perspective of the map).
func main() {
	lines := utils.ReadAndSplitRows("./test.txt", false, "")
	var startPositions []models.Coordinates
	linesToNumber := [][]int{}

	directions := []models.Coordinates{
		{X: 0, Y: -1}, // up
		{X: 0, Y: 1},  // down
		{X: -1, Y: 0}, // left
		{X: 1, Y: 0},  // right
	}

	for y, line := range lines {
		linesToNumber = append(linesToNumber, utils.ArrayAtoi(line))
		for x, char := range line {
			if char == string('0') {
				startPositions = append(startPositions, models.Coordinates{X: x, Y: y})
			}
		}
	}

	for _, start := range startPositions {
		// Process each starting position
		// Example: print the starting position
		println("Starting position:", start.X, start.Y)
	}

}
