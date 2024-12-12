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
				isAdiecentLeft := false
				isAdiecentRight := false
				isAdiecentUp := false
				isAdiecentDown := false
				//check prev
				if x > 0 && flowerMap[y][x-1].flowerName == flower.flowerName {
					isAdiecentLeft = true
				}
				// check next
				if x < len(flowerMap[y])-1 && flowerMap[y][x+1].flowerName == flower.flowerName {
					isAdiecentRight = true
				}
				//check up
				if y > 0 && flowerMap[y-1][x].flowerName == flower.flowerName {
					isAdiecentUp = true
				}
				//check down
				if y < len(flowerMap)-1 && flowerMap[y+1][x].flowerName == flower.flowerName {
					isAdiecentDown = true
				}

				if len(aiuole) == 0 {

					aiuole = append(aiuole, Aiuola{
						flowerType: flower.flowerName,
						perimeter:  getPerimiter(isAdiecentLeft, isAdiecentRight, isAdiecentUp, isAdiecentDown),
						area:       1,
						flowes:     []Flower{*flower},
					})
				} else {
					found := false
					for i := range aiuole {
						if aiuole[i].flowerType == flower.flowerName {
							for _, f := range aiuole[i].flowes {
								if (f.coordinates.X == flower.coordinates.X && (f.coordinates.Y == flower.coordinates.Y-1 || f.coordinates.Y == flower.coordinates.Y+1)) ||
									(f.coordinates.Y == flower.coordinates.Y && (f.coordinates.X == flower.coordinates.X-1 || f.coordinates.X == flower.coordinates.X+1)) {
									aiuole[i].perimeter += getPerimiter(isAdiecentLeft, isAdiecentRight, isAdiecentUp, isAdiecentDown)
									aiuole[i].area++
									aiuole[i].flowes = append(aiuole[i].flowes, *flower)
									found = true
									break
								}
							}
							break
						}
					}
					if !found {
						aiuole = append(aiuole, Aiuola{
							flowerType: flower.flowerName,
							perimeter:  getPerimiter(isAdiecentLeft, isAdiecentRight, isAdiecentUp, isAdiecentDown),
							area:       1,
							flowes:     []Flower{*flower},
						})
					}
				}

				flower.MarkVisited()
			}

		}
	}

	total := 0
	for _, aiuola := range aiuole {
		total += aiuola.perimeter * aiuola.area
	}
	fmt.Println("Total:", total)

}
