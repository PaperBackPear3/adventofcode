package main

import "adventofcode-2024/utils"

type Program struct {
	instruction []int
	value       []int
}

func main() {
	lines := utils.ReadFileLineByLine("./test.txt")

	registers := map[string]int{}

	program := Program{}

	for _, line := range lines {
		if line != "" {
			x := utils.SplitLineMultipleDelimiters(line, []string{"Register ", ": ", ","}, true)
			if x[0] != "Program" {
				registers[x[0]] = utils.Atoi(x[1])
			} else {
				for index, val := range x {
					if index == 0 {
						continue
					}
					if index%2 == 1 {
						program.instruction = append(program.instruction, utils.Atoi(val))
					} else {
						program.value = append(program.value, utils.Atoi(val))
					}
				}
			}
		}
	}
	println(program.instruction)
}

func part1() {
}
