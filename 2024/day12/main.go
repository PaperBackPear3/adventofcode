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
	part1(flowerMap)

	flowersGroup := map[string][]models.Coordinates{}

	for y, line := range lines {
		for x, char := range line {
			flowersGroup[char] = append(flowersGroup[char], models.Coordinates{X: x, Y: y})
		}
	}

	utils.PrintMapWithColors(lines, flowersGroup)
}

func part1(flowerMap [][]Flower) {
	aiuole := []Aiuola{}

	for y := range flowerMap {
		for x := range flowerMap[y] {
			flower := &flowerMap[y][x]
			if !flower.visited {
				if len(aiuole) == 0 {
					aiuole = append(aiuole, Aiuola{
						flowerType: flower.flowerName,
						flowes:     []Flower{*flower},
					})
				} else {
					found := false
					for i := range aiuole {
						if aiuole[i].flowerType == flower.flowerName {
							for _, f := range aiuole[i].flowes {
								if (f.coordinates.X == flower.coordinates.X && (f.coordinates.Y == flower.coordinates.Y-1 || f.coordinates.Y == flower.coordinates.Y+1)) ||
									(f.coordinates.Y == flower.coordinates.Y && (f.coordinates.X == flower.coordinates.X-1 || f.coordinates.X == flower.coordinates.X+1)) {
									aiuole[i].flowes = append(aiuole[i].flowes, *flower)
									found = true
									break
								}
							}
						}
					}
					if !found {
						aiuole = append(aiuole, Aiuola{
							flowerType: flower.flowerName,
							flowes:     []Flower{*flower},
						})
					}
				}
				flower.MarkVisited()
			}

		}
	}

	// Merge aiuole with at least one flower with the same coordinates
	for i := 0; i < len(aiuole); i++ {
		for j := i + 1; j < len(aiuole); j++ {
			merged := false
			for _, f1 := range aiuole[i].flowes {
				for _, f2 := range aiuole[j].flowes {
					if f1.coordinates == f2.coordinates {
						aiuole[i].flowes = append(aiuole[i].flowes, aiuole[j].flowes...)
						// Remove aiuole[j]
						aiuole = append(aiuole[:j], aiuole[j+1:]...)
						j--
						merged = true
						break
					}
				}
				if merged {
					break
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

}
