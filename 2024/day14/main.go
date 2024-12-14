package main

import (
	"adventofcode-2024/models"
	"adventofcode-2024/utils"
)

type Robot struct {
	startingCoordinates models.Coordinates
	velocity            models.Coordinates
}

func main() {
	lines := utils.ReadFileLineByLine("./test.txt")
	robotsMap := []Robot{}
	for _, line := range lines {
		robots := utils.SplitLineMultipleDelimiters(line, []string{" ", "p=", "v=", ","}, true)
		robotsMap = append(robotsMap, Robot{startingCoordinates: models.Coordinates{X: utils.Atoi(robots[0]), Y: utils.Atoi(robots[1])}, velocity: models.Coordinates{X: utils.Atoi(robots[2]), Y: utils.Atoi(robots[3])}})
	}
	println(robotsMap)

	grid := generateMap(robotsMap, 7, 11)
	utils.PrintMapWithColors(grid)
	part1(robotsMap, 7, 11)
}

func part1(robotsMap []Robot, mapHeight int, mapWidth int) {
}

func generateMap(robotsMap []Robot, mapHeight int, mapWidth int) [][]string {
	// Initialize the map with dots
	grid := make([][]string, mapHeight)
	for i := range grid {
		grid[i] = make([]string, mapWidth)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	// Place robots on the map
	for _, robot := range robotsMap {
		x, y := robot.startingCoordinates.X, robot.startingCoordinates.Y
		if x >= 0 && x < mapWidth && y >= 0 && y < mapHeight {
			if grid[y][x] == "." {
				grid[y][x] = "1"
			} else {
				count := utils.Atoi(grid[y][x])
				grid[y][x] = utils.Itoa(count + 1)
			}
		}
	}

	return grid
}
