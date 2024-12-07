package main

import "adventofcode-2024/utils"

type Row struct {
	number   int
	formedBy []int
}

func main() {
	lines := utils.ReadAndSplitRows("./input.txt", false, ":")
	SplittedLines := []Row{}
	for _, line := range lines {
		formedBy := utils.ArrayAtoi(utils.SplitStringToArray(line[1], true, " "))
		SplittedLines = append(SplittedLines, Row{number: utils.Atoi(line[0]), formedBy: formedBy})
	}
	println(SplittedLines)
	part1(SplittedLines)

}
func part1(lines []Row) {
	sum := 0
	validSolutions := [][]string{}

	for _, line := range lines {
		operatosCombo := generateOperatorCombinations(len(line.formedBy) - 1)
		for _, combo := range operatosCombo {
			if line.number == applyCombos(combo, line.formedBy) {
				validSolutions = append(validSolutions, combo)
				sum = sum + line.number
				break
			}
		}
	}
	println(validSolutions)
	println(sum)
}

func applyCombos(combos []string, values []int) int {
	result := values[0]
	for i := 1; i < len(values); i++ {
		if combos[i-1] == "+" {
			result = result + values[i]
		}
		if combos[i-1] == "*" {
			result = result * values[i]
		}
	}
	return result
}

func generateOperatorCombinations(n int) [][]string {
	// Generate all combinations of "+" and "*" for n positions without duplicates
	combinations := [][]string{}
	ops := []string{"+", "*"}

	// Helper function for recursive combination generation
	var helper func(current []string, depth int)
	helper = func(current []string, depth int) {
		if depth == n {
			// Add a copy of the current combination
			combination := make([]string, len(current))
			copy(combination, current)
			combinations = append(combinations, combination)
			return
		}
		for _, op := range ops {
			helper(append(current, op), depth+1)
		}
	}
	helper([]string{}, 0)
	return combinations
}
