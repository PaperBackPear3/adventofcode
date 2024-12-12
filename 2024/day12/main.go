package main

import (
	"adventofcode-2024/models"
	"adventofcode-2024/utils"
	"fmt"
)

type Flower struct {
	flowerName  string
	visited     bool
	coordinates models.Coordinates
}

// Method to mark the flower as visited
func (f *Flower) MarkVisited() {
	f.visited = true
}

// Method to mark the flower as visited
func getPerimiter(isAdiecentLeft bool, isAdiecentRight bool, isAdiecentUp bool, isAdiecentDown bool) int {
	perimeter := 4
	if isAdiecentLeft {
		perimeter--
	}
	if isAdiecentRight {
		perimeter--
	}
	if isAdiecentUp {
		perimeter--
	}
	if isAdiecentDown {
		perimeter--
	}
	return perimeter
}

type Aiuola struct {
	flowerType string
	perimeter  int
	area       int
	flowes     []Flower
}

func main() {

	lines := utils.ReadAndSplitRows("./test.txt", false, "")

	flowerMap := make([][]Flower, len(lines))
	for i := range flowerMap {
		flowerMap[i] = make([]Flower, len(lines[0]))
	}

	for y := range lines {
		for x := range lines[y] {
			flowerMap[y][x] = Flower{
				flowerName:  string(lines[y][x]),
				visited:     false,
				coordinates: models.Coordinates{X: x, Y: y},
			}
		}
	}
	aiuole := part1(flowerMap)
	part2(aiuole)

	flowersGroup := map[string][]models.Coordinates{}

	for y, line := range lines {
		for x, char := range line {
			flowersGroup[char] = append(flowersGroup[char], models.Coordinates{X: x, Y: y})
		}
	}

	utils.PrintMapWithColors(lines, flowersGroup)
}

func part1(flowerMap [][]Flower) []Aiuola {
	aiuole := []Aiuola{}

	var floodFill func(x, y int, flowerName string) []Flower
	floodFill = func(x, y int, flowerName string) []Flower {
		if x < 0 || x >= len(flowerMap[0]) || y < 0 || y >= len(flowerMap) {
			return nil
		}
		if flowerMap[y][x].visited || flowerMap[y][x].flowerName != flowerName {
			return nil
		}
		flowerMap[y][x].MarkVisited()
		flowers := []Flower{flowerMap[y][x]}
		flowers = append(flowers, floodFill(x-1, y, flowerName)...)
		flowers = append(flowers, floodFill(x+1, y, flowerName)...)
		flowers = append(flowers, floodFill(x, y-1, flowerName)...)
		flowers = append(flowers, floodFill(x, y+1, flowerName)...)
		return flowers
	}

	for y := range flowerMap {
		for x := range flowerMap[y] {
			if !flowerMap[y][x].visited {
				flowers := floodFill(x, y, flowerMap[y][x].flowerName)
				if len(flowers) > 0 {
					aiuole = append(aiuole, Aiuola{
						flowerType: flowerMap[y][x].flowerName,
						flowes:     flowers,
					})
				}
			}
		}
	}

	for i := range aiuole {
		aiuola := &aiuole[i]
		aiuola.area = len(aiuola.flowes)
		perimeter := 0
		for _, flower := range aiuola.flowes {
			x, y := flower.coordinates.X, flower.coordinates.Y
			isAdiecentLeft := x > 0 && flowerMap[y][x-1].flowerName == flower.flowerName
			isAdiecentRight := x < len(flowerMap[0])-1 && flowerMap[y][x+1].flowerName == flower.flowerName
			isAdiecentUp := y > 0 && flowerMap[y-1][x].flowerName == flower.flowerName
			isAdiecentDown := y < len(flowerMap)-1 && flowerMap[y+1][x].flowerName == flower.flowerName
			perimeter += getPerimiter(isAdiecentLeft, isAdiecentRight, isAdiecentUp, isAdiecentDown)
		}
		aiuola.perimeter = perimeter
	}

	total := 0
	for _, aiuola := range aiuole {
		total += aiuola.perimeter * aiuola.area
	}
	fmt.Println("Total:", total)
	return aiuole
}

func part2(aiuole []Aiuola) {
	// for _, aiuola := range aiuole {
	// 	for _, fiore := range aiuola.flowes {

	// 	}
	// }

}
