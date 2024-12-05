package utils

import (
	"math"
)

func SortArray(array []int) []int {
	// Sorts an array of elements
	// Returns the sorted array
	// Implement the sorting algorithm
	// Return the sorted array

	arrayLength := len(array)
	for i := 0; i < arrayLength; i++ {
		for j := 0; j < arrayLength; j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}

func ArrayAtoi(array []string) []int {
	// Convert an array of strings to an array of integers
	// Returns the array of integers
	// Implement the conversion
	// Return the array of integers
	intArray := make([]int, 0)
	for _, s := range array {
		intArray = append(intArray, Atoi(s))
	}
	return intArray
}

func ArrayValuesDiffAndMerge(array1 []int, array2 []int) []int {
	// Returns the difference of the values at the same index in two arrays
	// if one finishes before the other, the remaining values are returned the same

	diff := make([]int, 0)
	len1 := len(array1)
	len2 := len(array2)
	maxLen := len1
	if len2 > len1 {
		maxLen = len2
	}
	for i := 0; i < maxLen; i++ {
		if i < len1 && i < len2 {
			diff = append(diff, int(math.Abs(float64(array1[i]-array2[i]))))
		} else if i < len1 {
			diff = append(diff, array1[i])
		} else {
			diff = append(diff, array2[i])
		}
	}
	return diff
}

func ArraySum(array []int) int {
	sum := 0
	for _, val := range array {
		sum = sum + val
	}
	return sum
}

func ArrayFindOccasions(array []int, number int) int {
	occasions := 0
	for _, val := range array {
		if val == number {
			occasions++
		}
	}
	return occasions
}

func ArrayMultiply(array []int) int {

	product := 1
	for _, val := range array {
		product *= val
	}
	return product

}

func ArrayHas(array []int, value int) bool {
	for _, val := range array {
		if val == value {
			return true
		}
	}

	return false
}
