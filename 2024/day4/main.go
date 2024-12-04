package main

import (
	"adventofcode-2024/utils"
)

func main() {
	lines := utils.ReadAndSplitRows("./input.txt", false, "")
	height := len(lines)
	width := len(lines[0])
	coordinates := getCoordinates(lines, 'X')
	count := 0

	toMatch := "XMAS"
	reversedToMatch := "SAMX"
	for _, coord := range coordinates {
		internalCount := 0
		println("---CHECKING--- UP", coord.X, " - ", coord.Y)
		//check horizontal

		count, internalCount = checkHorizontalMatch(coord, width, lines, toMatch, count, internalCount)

		//check horizontal backwords
		if coord.X >= 3 {
			count, internalCount = checkHorizontalReverseMatch(lines, coord, reversedToMatch, count, internalCount)
		}
		// check positive oblique DOWN
		if coord.X < width-3 && coord.Y < height-3 {
			count, internalCount = checkPositiveDiagonalDownMatch(lines, coord, toMatch, count, internalCount)
		}
		// check positive oblique UP
		if coord.X < width-3 && coord.Y >= 3 {
			count, internalCount = checkPositiveDiagonalUpMatch(lines, coord, toMatch, count, internalCount)
		}
		// check backwards oblique DOWN
		if coord.X >= 3 && coord.Y < height-3 {
			count, internalCount = checkBackwardsObliqueDownMatch(lines, coord, toMatch, count, internalCount)
		}
		// check backwards oblique UP
		if coord.X >= 3 && coord.Y >= 3 {
			count, internalCount = checkBackwardsDiagonalUpMatch(lines, coord, toMatch, count, internalCount)
		}
		// check vertical down
		if coord.Y < height-3 {
			count, internalCount = checkVerticalMatchDown(lines, coord, toMatch, count, internalCount)
		}
		// check vertical up
		if coord.Y >= 3 {
			count, internalCount = checkVerticalMatchUp(lines, coord, toMatch, count, internalCount)

		}
		println("MATCH FOUND: ", internalCount)
		println("FINISHED CHECKS", coord.X, " - ", coord.Y)
	}
	println("count")
	println(count)
	part2(lines)
}

func checkVerticalMatchUp(lines [][]string, coord struct {
	X int
	Y int
}, toMatch string, count int, internalCount int) (int, int) {
	subset := ""
	for i := 0; i < 4; i++ {
		subset += string(lines[coord.Y-i][coord.X])
	}
	if subset == toMatch {
		count = count + 1
		internalCount = internalCount + 1
	}
	println("vertical UP", subset)
	return count, internalCount
}

func checkVerticalMatchDown(lines [][]string, coord struct {
	X int
	Y int
}, toMatch string, count int, internalCount int) (int, int) {
	subset := ""
	for i := 0; i < 4; i++ {
		subset += string(lines[coord.Y+i][coord.X])
	}
	if subset == toMatch {
		count = count + 1
		internalCount = internalCount + 1
	}
	println("vertical DOWN", subset)
	return count, internalCount
}

func checkBackwardsDiagonalUpMatch(lines [][]string, coord struct {
	X int
	Y int
}, toMatch string, count int, internalCount int) (int, int) {
	subset := ""
	for i := 0; i < 4; i++ {
		subset += string(lines[coord.Y-i][coord.X-i])
	}
	if subset == toMatch {
		count = count + 1
		internalCount = internalCount + 1
	}
	println("backwards oblique UP", subset)
	return count, internalCount
}

func checkBackwardsObliqueDownMatch(lines [][]string, coord struct {
	X int
	Y int
}, toMatch string, count int, internalCount int) (int, int) {
	subset := ""
	for i := 0; i < 4; i++ {
		subset += string(lines[coord.Y+i][coord.X-i])
	}
	if subset == toMatch {
		count = count + 1
		internalCount = internalCount + 1
	}
	println("backwards oblique DOWN", subset)
	return count, internalCount
}

func checkPositiveDiagonalUpMatch(lines [][]string, coord struct {
	X int
	Y int
}, toMatch string, count int, internalCount int) (int, int) {
	subset := ""
	for i := 0; i < 4; i++ {
		subset += string(lines[coord.Y-i][coord.X+i])
	}
	if subset == toMatch {
		count = count + 1
		internalCount = internalCount + 1
	}
	println("obllique UP", subset)
	return count, internalCount
}

func checkPositiveDiagonalDownMatch(lines [][]string, coord struct {
	X int
	Y int
}, toMatch string, count int, internalCount int) (int, int) {
	subset := ""
	for i := 0; i < 4; i++ {
		subset += string(lines[coord.Y+i][coord.X+i])
	}
	if subset == toMatch {
		count = count + 1
		internalCount = internalCount + 1
	}
	println("obllique DOWN", subset)
	return count, internalCount
}

func checkHorizontalReverseMatch(lines [][]string, coord struct {
	X int
	Y int
}, reversedToMatch string, count int, internalCount int) (int, int) {
	subset := lines[coord.Y][coord.X-3 : coord.X+1]
	subsetToString := ""
	for _, ch := range subset {
		subsetToString += string(ch)
	}

	if subsetToString == reversedToMatch {
		count = count + 1
		internalCount = internalCount + 1
	}
	println("HORIZONTAL BACKWORDS Subset:", subsetToString)
	return count, internalCount
}

func checkHorizontalMatch(coord struct {
	X int
	Y int
}, width int, lines [][]string, toMatch string, count int, internalCount int) (int, int) {
	if coord.X < width-3 {
		subset := lines[coord.Y][coord.X : coord.X+4]
		subsetToString := ""
		for _, ch := range subset {
			subsetToString += string(ch)
		}
		if subsetToString == toMatch {
			count = count + 1
			internalCount = internalCount + 1
		}
		println("HORIZONTAL Subset:", subsetToString)
	}
	return count, internalCount
}

func getCoordinates(lines [][]string, toCheck rune) []struct{ X, Y int } {
	var coordinates []struct{ X, Y int }

	for y, line := range lines {
		for x, char := range line {
			if len(char) == 1 && rune(char[0]) == toCheck {
				coordinates = append(coordinates, struct{ X, Y int }{X: x, Y: y})
			}
		}
	}
	return coordinates
}

func part2(lines [][]string) {
	coordinates := getCoordinates(lines, 'A')
	count := 0
	for _, coord := range coordinates {
		if coord.Y-1 >= 0 && coord.Y+1 < len(lines) && coord.X-1 >= 0 && coord.X+1 < len(lines[0]) {
			if lines[coord.Y-1][coord.X-1] == "M" || lines[coord.Y-1][coord.X-1] == "S" {
				if lines[coord.Y-1][coord.X+1] == "M" || lines[coord.Y-1][coord.X+1] == "S" {
					if lines[coord.Y+1][coord.X-1] == "M" || lines[coord.Y+1][coord.X-1] == "S" {
						if lines[coord.Y+1][coord.X+1] == "M" || lines[coord.Y+1][coord.X+1] == "S" {
							//check equity
							if lines[coord.Y-1][coord.X-1] != lines[coord.Y+1][coord.X+1] {
								if lines[coord.Y+1][coord.X-1] != lines[coord.Y-1][coord.X+1] {
									count = count + 1
								}

							}
						}
					}
				}
			}
		}
	}
	println(count)
}
