package main

import (
	"adventofcode-2024/utils"
	"math"
)

func main() {
	line := utils.ReadAndSplitRows("./input.txt", true, " ")
	count := 0
	toCheck := make([][]int, 0)
	for _, values := range line {
		toNumber := utils.ArrayAtoi(values)
		if allDecreasing(toNumber) {
			count++
		} else if allIncreasing(toNumber) {
			count++
		} else {
			toCheck = append(toCheck, toNumber)
		}
	}
	println(count)
	println(toCheck)
	for _, values := range toCheck {
		if isSafeWithOneRemoval(values) {
			count++
		}
	}
	println(count)
}

func isSafeWithOneRemoval(values []int) bool {
	arrCopy := make([]int, len(values))
	for i := range values {
		copy(arrCopy, values)
		// Create a new slice without the i-th element
		newValues := append(arrCopy[:i], arrCopy[i+1:]...)
		if allDecreasing(newValues) || allIncreasing(newValues) {
			return true
		}
	}
	return false
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
