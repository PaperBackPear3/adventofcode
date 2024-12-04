package main

import (
	"adventofcode-2024/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	rows := utils.ReadFileLineByLine("./input.txt")

	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	totale := 0
	for _, row := range rows {
		matches := pattern.FindAllStringSubmatch(row, -1)
		for _, match := range matches {
			fmt.Printf("Found pattern: %s\n", match[0])
			cleaned := regexp.MustCompile(`[^0-9,]`).ReplaceAllString(match[0], "")
			vals := utils.ArrayAtoi(strings.Split(cleaned, ","))
			totale = totale + utils.ArrayMultiply(vals)

		}
	}
	println(rows)
	println(totale)
	part2()
}

func part2() {
	rows := utils.ReadFileLineByLine("./input.txt")
	totale := 0
	for index, row := range rows {
		if strings.Contains(row, "don't()") {
			parts := strings.Split(row, "don't()")

			for _, donts := range parts {
				if strings.Contains(donts, "do()") || index == 0 {
					dos := strings.Split(donts, "do()")
					startLine := 0
					if index != 0 {
						startLine = 1
					}
					for _, dosPart := range dos[startLine:] {
						pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
						matches := pattern.FindAllStringSubmatch(dosPart, -1)
						for _, match := range matches {
							cleaned := regexp.MustCompile(`[^0-9,]`).ReplaceAllString(match[0], "")
							vals := utils.ArrayAtoi(strings.Split(cleaned, ","))
							totale = totale + utils.ArrayMultiply(vals)
						}
					}
				}
			}

			// pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
			// for _, part := range parts {
			// 	matches := pattern.FindAllStringSubmatch(part, -1)
			// 	for i := 1; i < len(matches); i++ {
			// 		match := matches[i]
			// 		cleaned := regexp.MustCompile(`[^0-9,]`).ReplaceAllString(match[0], "")
			// 		vals := utils.ArrayAtoi(strings.Split(cleaned, ","))
			// 		totale = totale + utils.ArrayMultiply(vals)

			// 	}
			// }
		}
	}
	println(rows)
	println(totale)

}
