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
	lines := utils.ReadAndSplitRows("./input.txt", false, "")
	part1(lines)
}

func part1(lines [][]string) {
	guard := findGuards(lines)[0]
	fmt.Printf("Guard: %+v\n", guard)
	maxX := len(lines[0])
	maxY := len(lines)
	stepTaken := 1
	for {
		if !isWithinBoundaries(guard.X, guard.Y, maxX, maxY, guard.Direction) {
			break
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
	}
	println(stepTaken)
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

func isWithinBoundaries(x, y, maxX, maxY int, direction Direction) bool {
	switch direction {
	case Right:
		return x >= 0 && x < maxX-1 && y >= 0 && y < maxY
	case Left:
		return x > 0 && x < maxX && y >= 0 && y < maxY
	case Up:
		return x >= 0 && x < maxX && y > 0 && y < maxY
	case Down:
		return x >= 0 && x < maxX && y >= 0 && y < maxY-1
	default:
		return false
	}
}
