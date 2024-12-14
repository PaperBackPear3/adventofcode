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
	mapWidth := 11 //101
	mapHeight := 7 //103

	grid, robotsLocations := generateMap(robotsMap, mapHeight, mapWidth)
	utils.PrintMapWithColors(grid, robotsLocations)
	part1(robotsMap, mapHeight, mapWidth)
}

func part1(robotsMap []Robot, mapHeight int, mapWidth int) {
}

func generateMap(robotsMap []Robot, mapHeight int, mapWidth int) ([][]string, map[string][]models.Coordinates) {
	// Initialize the map with dots
	grid := make([][]string, mapHeight)
	for i := range grid {
		grid[i] = make([]string, mapWidth)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	// Initialize the robots map
	robotsCoordinates := make(map[string][]models.Coordinates)

	// Place robots on the map
	for index, robot := range robotsMap {
		x, y := robot.startingCoordinates.X, robot.startingCoordinates.Y
		if x >= 0 && x < mapWidth && y >= 0 && y < mapHeight {
			if grid[y][x] == "." {
				grid[y][x] = "1"
			} else {
				count := utils.Atoi(grid[y][x])
				grid[y][x] = utils.Itoa(count + 1)
			}
			robotsCoordinates[utils.Itoa(index)] = append(robotsCoordinates[utils.Itoa(index)], models.Coordinates{X: x, Y: y})
		}
	}

	return grid, robotsCoordinates
}
