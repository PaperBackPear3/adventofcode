package utils

import (
	"log"
	"strconv"
	"strings"
)

func SplitStringToArray(line string, removeWhiteSpaces bool, delimiter string) []string {
	// Split by columns
	// Returns an array with all the columns values
	splittedLine := strings.Split(line, delimiter)
	if removeWhiteSpaces {
		toAdd := make([]string, 0)
		//add non empty columns to the array
		for _, column := range splittedLine {
			if column != " " && column != "" {
				toAdd = append(toAdd, column)
			}
		}
		splittedLine = toAdd
	}
	return splittedLine

}

func Atoi(s string) int {
	// Convert a string to an integer
	// Returns the integer
	// Implement the conversion
	// Return the integer
	i, err := strconv.Atoi(s)

	if err != nil {
		log.Panic(err)
	}
	return i
}

func SplitLineMultipleDelimiters(line string, delimitersList []string, removeWhiteSpaces bool) []string {
	if len(delimitersList) == 0 {
		return []string{line}
	}

	// Split the line by the first delimiter
	parts := strings.Split(line, delimitersList[0])

	// If there are more delimiters, split each part further
	if len(delimitersList) > 1 {
		var result []string
		for _, part := range parts {
			result = append(result, SplitLineMultipleDelimiters(part, delimitersList[1:], removeWhiteSpaces)...)
		}
		parts = result
	}

	if removeWhiteSpaces {
		toAdd := make([]string, 0)
		for _, part := range parts {
			if part != " " && part != "" {
				toAdd = append(toAdd, part)
			}
		}
		parts = toAdd
	}

	return parts
}
