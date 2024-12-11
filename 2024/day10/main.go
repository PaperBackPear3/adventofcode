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
	var endPositions []models.Coordinates
	linesToNumber := [][]int{}
	for y, line := range lines {
		linesToNumber = append(linesToNumber, utils.ArrayAtoi(line))
		for x, char := range line {
			if char == string('0') {
				startPositions = append(startPositions, models.Coordinates{X: x, Y: y})
			}
			if char == string('9') {
				endPositions = append(endPositions, models.Coordinates{X: x, Y: y})
			}
		}
	}

	count := 0
	for _, start := range startPositions {
		// Process each start position
		score := 0

		visited := make(map[models.Coordinates]bool)
		var findPath func(current models.Coordinates) bool
		findPath = func(current models.Coordinates) bool {
			if visited[current] {
				return false
			}
			visited[current] = true

			if currentValue := linesToNumber[current.Y][current.X]; currentValue == 9 {
				score = score + 1
				return true
			}

			directions := []models.Coordinates{
				{X: 0, Y: -1}, // up
				{X: 0, Y: 1},  // down
				{X: -1, Y: 0}, // left
				{X: 1, Y: 0},  // right
			}

			for _, direction := range directions {
				next := models.Coordinates{X: current.X + direction.X, Y: current.Y + direction.Y}
				if next.Y >= 0 && next.Y < len(linesToNumber) && next.X >= 0 && next.X < len(linesToNumber[next.Y]) {
					if linesToNumber[next.Y][next.X] == linesToNumber[current.Y][current.X]+1 {
						if findPath(next) {
							return true
						}
					}
				}
			}

			return false
		}

		if findPath(start) {
			count = count + 1
		}
		println(score)
	}

	print(count)

	println(startPositions)
	println(endPositions)
}
