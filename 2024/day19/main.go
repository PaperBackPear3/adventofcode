package main

import (
	"adventofcode-2024/utils"
)

func main() {

	lines := utils.ReadFileLineByLine("./input.txt")
	var availableTowels []string
	var designsToMake []string

	for index, line := range lines {
		if index == 0 {
			availableTowels = utils.SplitStringToArray(line, true, ", ")
			println(availableTowels)
		} else if line != "" {
			designsToMake = append(designsToMake, line)
		}
	}
	part1(designsToMake, availableTowels)
	part2(designsToMake, availableTowels)

}

func waysToMakeDesign(design string, towels []string, cache map[string]int) int {
	// Base case: If the design is empty, we've successfully matched it
	if design == "" {
		return 1
	}

	// Check if the result is already cached
	if val, exists := cache[design]; exists {
		return val
	}

	// Calculate the number of ways recursively
	totalWays := 0
	for _, towel := range towels {
		if len(design) >= len(towel) && design[:len(towel)] == towel {
			totalWays += waysToMakeDesign(design[len(towel):], towels, cache)
		}
	}

	// Cache the result and return it
	cache[design] = totalWays
	return totalWays
}

func part1(designs []string, towels []string) int {
	cache := make(map[string]int)
	count := 0

	for _, design := range designs {
		if waysToMakeDesign(design, towels, cache) > 0 {
			count++
		}
	}

	return count
}

func part2(designs []string, towels []string) int {
	cache := make(map[string]int)
	count := 0

	for _, design := range designs {
		count = count + waysToMakeDesign(design, towels, cache)
	}

	return count
}
