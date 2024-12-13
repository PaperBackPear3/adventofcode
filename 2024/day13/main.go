package main

import (
	"adventofcode-2024/models"
	"adventofcode-2024/utils"
)

type Machine struct {
	buttonA models.Coordinates
	buttonB models.Coordinates
	prize   models.Coordinates
}

func main() {

	lines := utils.ReadFileLineByLine("./test.txt")
	input := []Machine{}
	newMachine := Machine{}
	for lineIndex, line := range lines {
		if lineIndex == 0 {
			newMachine = Machine{}
		}
		if line != "" {
			tmp := utils.SplitLineMultipleDelimiters(line, []string{"Button ", ":", ",", "X+", "Y+", "X=", "Y="}, true)
			if tmp[0] == "A" {
				newMachine.buttonA.X = utils.Atoi(tmp[1])
				newMachine.buttonA.Y = utils.Atoi(tmp[2])
			}
			if tmp[0] == "B" {
				newMachine.buttonB.X = utils.Atoi(tmp[1])
				newMachine.buttonB.Y = utils.Atoi(tmp[2])
			}
			if tmp[0] == "Prize" {
				newMachine.prize.X = utils.Atoi(tmp[1])
				newMachine.prize.Y = utils.Atoi(tmp[2])
			}
		}
		if line == "" || lineIndex == len(lines)-1 {
			input = append(input, newMachine)
			newMachine = Machine{}
		}
	}
	println(input)
	part1(input)
}

func part1(machines []Machine) {
}
