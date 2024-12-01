package main

import (
	"adventofcode-2024/utils"
)

func part2(filepath string) {
	test := utils.ReadAndSplitColumns(filepath, true, " ")
	integerArray := make([][]int, 0)
	for i, row := range test {
		if len(integerArray) < i+1 {
			integerArray = append(integerArray, make([]int, 0))
		}
		integerArray = append(integerArray, utils.ArrayAtoi(row))
		integerArray[i+1] = utils.SortArray(integerArray[i+1])
	}
	merged := utils.ArrayValuesDiffAndMerge(integerArray[1], integerArray[2])
	println(utils.ArraySum(merged))
}
