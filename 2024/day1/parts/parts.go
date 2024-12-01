package parts

import "adventofcode-2024/utils"

func Part1(filepath string) {
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

func Part2(filepath string) {
	test := utils.ReadAndSplitColumns(filepath, true, " ")
	integerArray := make([][]int, 0)
	for i, row := range test {
		if len(integerArray) < i+1 {
			integerArray = append(integerArray, make([]int, 0))
		}
		integerArray = append(integerArray, utils.ArrayAtoi(row))
		integerArray[i+1] = utils.SortArray(integerArray[i+1])
	}

	final := make([]int, 0)

	for _, val := range integerArray[1] {
		occasions := utils.ArrayFindOccasions(integerArray[2], val)
		if occasions > 0 {
			final = append(final, val*occasions)
		}
	}
	println(utils.ArraySum(final))
}
