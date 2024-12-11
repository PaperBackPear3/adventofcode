package main

import (
	"adventofcode-2024/utils"
	"math"
	"strconv"
)

func digitCount(n int) int {
	if n == 0 {
		return 1
	}
	// Handle negative numbers
	n = int(math.Abs(float64(n)))
	count := 0
	for n > 0 {
		count++
		n /= 10
	}
	return count
}

func main() {
	line := utils.ReadAndSplitRows("./input.txt", false, " ")
	numbers := utils.ArrayAtoi(line[0])
	println(numbers)
	for i := 0; i < 25; i++ {
		newNumbers := make([]int, 0, len(numbers)*2) // Preallocate with an estimated size

		for _, number := range numbers {
			if number == 0 {
				newNumbers = append(newNumbers, 1)
			} else if digitCount(number)%2 == 0 {
				strNum := strconv.Itoa(number)
				mid := len(strNum) / 2
				left, _ := strconv.Atoi(strNum[:mid])
				right, _ := strconv.Atoi(strNum[mid:])
				newNumbers = append(newNumbers, left, right)
			} else {
				newNumbers = append(newNumbers, number*2024)
			}
		}
		numbers = newNumbers
	}

	println(len(numbers))
}
