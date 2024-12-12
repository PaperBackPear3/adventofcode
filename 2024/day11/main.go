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
	firstLine := line[0]

	part1(firstLine, 25)
	part2(firstLine)
}

func part1(firstLine []string, loops int) int {
	numbers := utils.ArrayAtoi(firstLine)
	for i := 0; i < loops; i++ {
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
	return len(numbers)
}

/// part2

func part2(firstLine []string) {
	println("Part2: ")
	numbers := utils.ArrayAtoi(firstLine)
	testZ := []string{"0"}
	zerosLength := part1(testZ, 75)
	zeros := 0
	ones := 0
	println(numbers)
	for i := 0; i < 75; i++ {
		newNumbers := make([]int, 0, len(numbers)*2) // Preallocate with an estimated size

		for _, number := range numbers {
			if number == 0 {
				zeros = zeros + 1
			} else if number == 1 {
				ones = ones + 1
			} else if digitCount(number)%2 == 0 {
				strNum := strconv.Itoa(number)
				mid := len(strNum) / 2
				left, _ := strconv.Atoi(strNum[:mid])
				right, _ := strconv.Atoi(strNum[mid:])
				if left == 0 {
					zeros = zeros + 1
					newNumbers = append(newNumbers, right)
				} else if right == 0 {
					zeros = zeros + 1
					newNumbers = append(newNumbers, left)
				} else {
					newNumbers = append(newNumbers, left, right)
				}
			} else {
				newNumbers = append(newNumbers, number*2024)
			}
		}
		numbers = newNumbers
	}

	println(len(numbers))
	println(zerosLength)
	println(zeros * zerosLength)
	println(ones * (zerosLength - 1))
}
