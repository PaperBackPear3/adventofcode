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

	lines := utils.ReadFileLineByLine("./input.txt")
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
				newMachine.prize.X = utils.Atoi(tmp[1]) * 1000000000
				newMachine.prize.Y = utils.Atoi(tmp[2]) * 1000000000
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

	type ButtonPressed struct {
		ButtonAPressedCount int
		ButtonBPressedCount int
	}
	allWorkingCombos := []ButtonPressed{}

	type AllMachinesList struct {
		machineId        int
		allWorkingCombos []ButtonPressed
	}

	allMachines := []AllMachinesList{}

	for machineID, test := range machines {
		allWorkingCombos = []ButtonPressed{}
		prizeX := test.prize.X
		prizeY := test.prize.Y

		buttonAPress := 0
		buttonBPress := 0

		currentX := 0
		currentY := 0

		for {
			for {
				if currentX < prizeX && currentY < prizeY {
					buttonBPress = buttonBPress + 1
					currentX = currentX + test.buttonB.X
					currentY = currentY + test.buttonB.Y
				} else if currentX == prizeX && currentY == prizeY {
					allWorkingCombos = append(allWorkingCombos, ButtonPressed{ButtonAPressedCount: buttonAPress, ButtonBPressedCount: buttonBPress})
					buttonBPress = 0
					currentX = buttonAPress * test.buttonA.X
					currentY = buttonAPress * test.buttonA.Y
					break
				} else {
					buttonBPress = 0
					currentX = buttonAPress * test.buttonA.X
					currentY = buttonAPress * test.buttonA.Y
					break
				}
			}
			if currentX < prizeX && currentY < prizeY {
				buttonAPress = buttonAPress + 1
				currentX = currentX + test.buttonA.X
				currentY = currentY + test.buttonA.Y
			} else if currentX == prizeX && currentY == prizeY {
				allWorkingCombos = append(allWorkingCombos, ButtonPressed{ButtonAPressedCount: buttonAPress, ButtonBPressedCount: buttonBPress})
				buttonAPress = 0
				break
			} else {
				buttonAPress = 0
				break
			}
		}
		if len(allWorkingCombos) > 0 {
			allMachines = append(allMachines, AllMachinesList{
				machineId:        machineID,
				allWorkingCombos: allWorkingCombos,
			})
		}
	}
	globalCost := 0
	for _, machine := range allMachines {
		for _, combo := range machine.allWorkingCombos {
			cost := combo.ButtonAPressedCount*3 + combo.ButtonBPressedCount
			globalCost = globalCost + cost
		}
	}
	println("global cost: ", globalCost)
}
