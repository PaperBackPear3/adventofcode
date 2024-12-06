package main

import (
	"adventofcode-2024/utils"
	"fmt"
)

type Direction string

const (
	Right Direction = ">"
	Left  Direction = "<"
	Up    Direction = "^"
	Down  Direction = "v"
)

var AllDirections = []Direction{Right, Left, Up, Down}

type Guard struct {
	Direction Direction
	X         int
	Y         int
}

func main() {
	lines := utils.ReadAndSplitRows("./test.txt", false, "")
	linesCopy := make([][]string, len(lines))
	for i := range lines {
		linesCopy[i] = make([]string, len(lines[i]))
		copy(linesCopy[i], lines[i])
	}
	guard := findGuards(lines)[0]
	lines[guard.Y][guard.X] = string(guard.Direction)

	part1(lines, guard, -1, -1)
	lines[guard.Y][guard.X] = string(guard.Direction)
	println("first map:")
	printMap(lines, guard)
	part2(lines, guard)
	println("final map:")
	printMap(lines, guard)
}

func part1(lines [][]string, guard Guard, fakeObstacleX int, fakeObstacleY int) bool {
	fmt.Printf("Guard: %+v\n", guard)
	maxX := len(lines[0])
	maxY := len(lines)
	stepTaken := 0
	loop := 0
	for {
		if loop < 5 {
			if isAtFakeObstacle(guard, fakeObstacleX, fakeObstacleY) {
				loop++
			}
		} else {
			fmt.Println("loop detected", fakeObstacleX, fakeObstacleY)
			return false
		}

		switch guard.Direction {
		case Right:
			if lines[guard.Y][guard.X+1] == "#" {
				guard.Direction = Down
			} else {
				if lines[guard.Y][guard.X+1] != "X" {
					stepTaken++
				}
				lines[guard.Y][guard.X] = "X"
				guard.X++
			}
		case Left:
			if lines[guard.Y][guard.X-1] == "#" {
				guard.Direction = Up
			} else {
				if lines[guard.Y][guard.X-1] != "X" {
					stepTaken++
				}
				lines[guard.Y][guard.X] = "X"
				guard.X--
			}
		case Up:
			if lines[guard.Y-1][guard.X] == "#" {
				guard.Direction = Right
			} else {
				if lines[guard.Y-1][guard.X] != "X" {
					stepTaken++
				}
				lines[guard.Y][guard.X] = "X"
				guard.Y--
			}
		case Down:
			if lines[guard.Y+1][guard.X] == "#" {
				guard.Direction = Left
			} else {
				if lines[guard.Y+1][guard.X] != "X" {
					stepTaken++
				}
				lines[guard.Y][guard.X] = "X"
				guard.Y++
			}
		}
		if !isWithinBoundaries(guard.X, guard.Y, maxX, maxY, guard.Direction) {
			break
		}
	}
	println("steps: ", stepTaken)
	return true
}

func findGuards(lines [][]string) []Guard {
	var guards []Guard
	for y, line := range lines {
		for x, char := range line {
			for _, direction := range AllDirections {
				if string(char) == string(direction) {
					guards = append(guards, Guard{Direction: direction, X: x, Y: y})
				}
			}
		}
	}
	return guards
}

func isAtFakeObstacle(guard Guard, fakeObstacleX int, fakeObstacleY int) bool {
	switch guard.Direction {
	case Right:
		return guard.X+1 == fakeObstacleX && guard.Y == fakeObstacleY
	case Left:
		return guard.X-1 == fakeObstacleX && guard.Y == fakeObstacleY
	case Up:
		return guard.X == fakeObstacleX && guard.Y-1 == fakeObstacleY
	case Down:
		return guard.X == fakeObstacleX && guard.Y+1 == fakeObstacleY
	default:
		return false
	}
}

func isWithinBoundaries(x int, y int, maxX int, maxY int, direction Direction) bool {
	switch direction {
	case Right:
		return x+1 < maxX
	case Left:
		return x-1 >= 0
	case Up:
		return y-1 >= 0
	case Down:
		return y+1 < maxY
	default:
		return false
	}
}

func part2(linesWithPath [][]string, guard Guard) {
	fmt.Printf("Guard: %+v\n", guard)
	count := 0
	for y := range linesWithPath {
		for x := range linesWithPath[y] {
			if linesWithPath[y][x] == "X" {
				linesCopy := make([][]string, len(linesWithPath))
				for i := range linesWithPath {
					linesCopy[i] = make([]string, len(linesWithPath[i]))
					copy(linesCopy[i], linesWithPath[i])
				}
				linesCopy[y][x] = "#"
				if !part1(linesCopy, guard, x, y) {
					linesWithPath[y][x] = "O"
					count = count + 1
					printMap(linesWithPath, guard)
				}
			}
		}
	}
	println("count: ", count)
}

func printMap(lines [][]string, guard Guard) {
	for y, row := range lines {
		for x, cell := range row {
			if x == guard.X && y == guard.Y {
				fmt.Printf("\033[34m%s\033[0m", guard.Direction) // Print guard position in blue
			} else if cell == "X" {
				fmt.Print("\033[32mX\033[0m") // Print path in green
			} else if cell == "#" {
				fmt.Print("\033[31m#\033[0m") // Print obstacle in red
			} else if cell == "O" {
				fmt.Print("\033[33mO\033[0m") // Print marked cell in yellow
			} else {
				fmt.Print(cell)
			}
		}
		fmt.Println()
	}
}
