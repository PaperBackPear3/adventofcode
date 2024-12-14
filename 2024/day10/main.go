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

type Cell struct {
	value   int
	visited bool
}

func main() {
	lines := utils.ReadAndSplitRows("./test.txt", false, "")
	var startPositions []models.Coordinates

	worldMap := [][]Cell{}

	for y, line := range lines {
		row := []Cell{}
		for x, char := range line {
			number := utils.Atoi(char)
			row = append(row, Cell{value: number, visited: false})
			if char == "0" {
				startPositions = append(startPositions, models.Coordinates{X: x, Y: y})
			}
		}
		worldMap = append(worldMap, row)
	}
	part1(worldMap, startPositions)
}

//a valid path is one that starts at 0 and arrives at 9
//from all starting position floodFill until it finds a 9
// to move from one cell to the next one there must be a difference of at most 1. 0 can go to 1 and 1 to 2 not more.

func part1(worldMap [][]Cell, startPositions []models.Coordinates) {
	paths := make(map[models.Coordinates]int)
	for _, start := range startPositions {
	}

	for start, count := range paths {
		println("Starting position:", start.X, "Total paths to 9:", count)
	}
}
