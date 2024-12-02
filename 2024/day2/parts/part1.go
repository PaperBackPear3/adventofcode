package main

import (
	"adventofcode-2024/utils"
	"math"
)

func main() {
	line := utils.ReadAndSplitRows("./input.txt", true, " ")
	count := 0
	for _, values := range line {
		toNumber := utils.ArrayAtoi(values)
		if allDecreasing(toNumber) {
			count++
		}
		if allIncreasing(toNumber) {
			count++
		}
	}
	println(count)
}

func allDecreasing(values []int) bool {
	valid := true
	for i := 0; i < len(values)-1; i++ {
		//all decreasing.
		if values[i] >= values[i+1] && int(math.Abs(float64(values[i]-values[i+1]))) > 0 && int(math.Abs(float64(values[i]-values[i+1]))) <= 3 {
		} else {
			return false
		}
	}
	return valid
}

func allIncreasing(values []int) bool {
	valid := true
	for i := 0; i < len(values)-1; i++ {
		//all increasing.
		if values[i] <= values[i+1] && int(math.Abs(float64(values[i]-values[i+1]))) > 0 && int(math.Abs(float64(values[i]-values[i+1]))) <= 3 {
		} else {
			return false
		}
	}
	return valid
}
