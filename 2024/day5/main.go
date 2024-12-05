package main

import (
	"adventofcode-2024/utils"
	"log"
)

//first are ordering rules for pages
//second inputs are for pages to produce in each update

func main() {

	inputs := utils.ReadMultipleTypesOfInputs("test.txt", "")
	var pageNumbers [][]int
	for _, input := range inputs[1] {
		log.Println(input)

		pageNumbers = append(pageNumbers, utils.ArrayAtoi(utils.SplitStringToArray(input, false, ",")))
	}

	daRulez := map[int][]int{}

	for _, ruleLine := range inputs[0] {
		toNumbers := utils.ArrayAtoi(utils.SplitStringToArray(ruleLine, false, "|"))
		key := toNumbers[0]
		daRulez[key] = append(daRulez[key], toNumbers[1])
	}
	savedPages, wrongPages := part1(daRulez, pageNumbers)

	middlePages := []int{}
	for _, pages := range savedPages {
		middle := pages[len(pages)/2]
		middlePages = append(middlePages, middle)
	}
	println(utils.ArraySum(middlePages))

	fixed := part2(wrongPages)
	savedPages, _ = part1(daRulez, fixed)
	for _, pages := range savedPages {
		middle := pages[len(pages)/2]
		middlePages = append(middlePages, middle)
	}
	println(utils.ArraySum(middlePages))

}

func part1(rules map[int][]int, pageNumbers [][]int) ([][]int, [][]int) {

	savedPages := [][]int{}
	wrongPages := [][]int{}
	for _, pages := range pageNumbers {
		isCorrect := true
		for i := 1; i < len(pages); i++ {
			currentRule := rules[pages[i]]
			for j := 0; j < i; j++ {
				if utils.ArrayHas(currentRule, pages[j]) {
					isCorrect = false
					break
				}
			}

		}
		if isCorrect {
			savedPages = append(savedPages, pages)
		} else {
			wrongPages = append(wrongPages, pages)
		}
	}
	return savedPages, wrongPages
}

func part2(wrongPages [][]int) [][]int {
	fixedPages := [][]int{}
	return fixedPages
}
