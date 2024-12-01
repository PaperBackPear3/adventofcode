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
