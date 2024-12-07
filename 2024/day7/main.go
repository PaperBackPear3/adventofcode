package main

import (
	"adventofcode-2024/utils"
	"fmt"
	"strconv"
)

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
	toFixSolutions, sum1 := part1(SplittedLines)
	sum1 = sum1 + part2(toFixSolutions)
	println(sum1)
}
func part1(lines []Row) ([]Row, int) {
	sum := 0
	toFixSolutions := []Row{}

	for _, line := range lines {
		operatosCombo := generateOperatorCombinations(len(line.formedBy)-1, []string{"+", "*"})
		count := 0

		for _, combo := range operatosCombo {
			if line.number == applyCombos(combo, line.formedBy) {
				sum = sum + line.number
				count = count + 1
				break
			}
		}
		if count == 0 {
			toFixSolutions = append(toFixSolutions, line)
		}
	}
	println(sum)
	return toFixSolutions, sum
}

func part2(lines []Row) int {
	sum := 0

	for _, line := range lines {
		operatosCombo := generateOperatorCombinations(len(line.formedBy)-1, []string{"+", "*", "||"})
		for _, combo := range operatosCombo {
			if line.number == applyCombosPart2(combo, line.formedBy) {
				sum = sum + line.number
				break
			}
		}
	}
	return sum
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

func applyCombosPart2(combos []string, values []int) int {
	result := values[0]
	for i := 1; i < len(values); i++ {
		if combos[i-1] == "+" {
			result = result + values[i]
		}
		if combos[i-1] == "*" {
			result = result * values[i]
		}
		if combos[i-1] == "||" {
			resultStr := fmt.Sprintf("%d%d", result, values[i])
			result, _ = strconv.Atoi(resultStr)
		}
	}
	return result
}

func generateOperatorCombinations(n int, ops []string) [][]string {
	// Generate all combinations of "+" and "*" for n positions without duplicates
	combinations := [][]string{}

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
