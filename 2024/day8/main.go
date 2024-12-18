package main

import (
	"adventofcode-2024/models"
	"adventofcode-2024/utils"
)

func main() {
	lines := utils.ReadAndSplitRows("./input.txt", false, "")

	// contains the antenna type as key and an array of all the same types coordinates
	antennasLocations := map[string][]models.Coordinates{}

	for y, line := range lines {
		for x, char := range line {
			if char != "." {
				antennasLocations[char] = append(antennasLocations[char], models.Coordinates{X: x, Y: y})
			}
		}
	}
	// Make a copy of the lines
	linesCopy := make([][]string, len(lines))
	for i := range lines {
		linesCopy[i] = make([]string, len(lines[i]))
		copy(linesCopy[i], lines[i])
	}
	part1(lines, antennasLocations)
	part2(linesCopy, antennasLocations)
}

func part1(antennasMap [][]string, antennasLocation map[string][]models.Coordinates) {

	maxX := len(antennasMap[0])
	maxY := len(antennasMap)
	for _, locations := range antennasLocation {
		manhattanDistance := map[models.Coordinates]map[models.Coordinates]models.Coordinates{}
		for _, coords := range locations {
			for i := 0; i < len(locations); i++ {
				//no distance if the coordinates are the same
				if coords.X != locations[i].X && coords.Y != locations[i].Y {
					distX := (float64(coords.X) - float64(locations[i].X))
					distY := (float64(coords.Y) - float64(locations[i].Y))
					if manhattanDistance[coords] == nil {
						manhattanDistance[coords] = make(map[models.Coordinates]models.Coordinates)
					}
					manhattanDistance[coords][locations[i]] = models.Coordinates{X: int(distX), Y: int(distY)}
				}
			}
		}
		for index, antennaLocation := range manhattanDistance {
			toAddX := index.X
			toAddY := index.Y
			for _, otherDistances := range antennaLocation {
				newCoords := models.Coordinates{X: toAddX + otherDistances.X, Y: toAddY + otherDistances.Y}
				if newCoords.X < maxX && newCoords.Y < maxY && newCoords.X >= 0 && newCoords.Y >= 0 {
					antennasMap[newCoords.Y][newCoords.X] = "#"
				}
			}
		}
	}
	utils.PrintMapWithColors(antennasMap, antennasLocation)
	totalX := 0
	for _, row := range antennasMap {
		for _, cell := range row {
			if cell == "#" {
				totalX++
			}
		}
	}
	println("Total X count:", totalX)
}

func part2(antennasMap [][]string, antennasLocation map[string][]models.Coordinates) {

	maxX := len(antennasMap[0])
	maxY := len(antennasMap)
	for _, locations := range antennasLocation {
		manhattanDistance := map[models.Coordinates]map[models.Coordinates]models.Coordinates{}
		for _, coords := range locations {
			for i := 0; i < len(locations); i++ {
				//no distance if the coordinates are the same
				if coords.X != locations[i].X && coords.Y != locations[i].Y {
					distX := (float64(coords.X) - float64(locations[i].X))
					distY := (float64(coords.Y) - float64(locations[i].Y))
					if manhattanDistance[coords] == nil {
						manhattanDistance[coords] = make(map[models.Coordinates]models.Coordinates)
					}
					manhattanDistance[coords][locations[i]] = models.Coordinates{X: int(distX), Y: int(distY)}
				}
			}
		}

		for index, antennaLocation := range manhattanDistance {
			toAddX := index.X
			toAddY := index.Y
			for _, otherDistances := range antennaLocation {

				newCoords := models.Coordinates{X: toAddX, Y: toAddY}
				for {
					newCoords.X += otherDistances.X
					newCoords.Y += otherDistances.Y
					if newCoords.X < 0 || newCoords.Y < 0 || newCoords.X >= maxX || newCoords.Y >= maxY {
						break
					}
					antennasMap[newCoords.Y][newCoords.X] = "#"
				}
			}
		}
	}
	utils.PrintMapWithColors(antennasMap, antennasLocation)
	totalX := 0
	for _, row := range antennasMap {
		for _, cell := range row {
			if cell != "." {
				totalX++
			}
		}
	}
	println("Total X count:", totalX)
}
