package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadAndSplitColumns(filePath string, removeWhiteSpaces bool, delimiter string) [][]string {
	// Read file
	// Split by columns
	// Returns ad array of arrays with all the columns values
	file, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	// Read file line by line
	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)
	splittedLine := make([][]string, 0)
	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		if os.Getenv("DEBUG") == "true" {
			fmt.Printf("line: %s\n", line)
		}
		// Split by columns
		splitted := SplitStringToArray(line, removeWhiteSpaces, delimiter)
		//add each element of splitted to splitledline creating N columns based on splitted length
		for i, column := range splitted {
			//append to the Ith column of splittedLine the Ith element of splitted
			if len(splittedLine) < i+1 {
				splittedLine = append(splittedLine, make([]string, 0))
			}
			splittedLine[i] = append(splittedLine[i], column)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return splittedLine

}

func ReadAndSplitRows(filePath string, removeWhiteSpaces bool, delimiter string) [][]string {
	// Read file
	// Split by columns
	// Returns ad array of arrays with all the columns values
	file, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	// Read file line by line
	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)
	splittedLine := make([][]string, 0)
	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		if os.Getenv("DEBUG") == "true" {
			fmt.Printf("line: %s\n", line)
		}
		// Split by columns
		splitted := SplitStringToArray(line, removeWhiteSpaces, delimiter)
		//add each element of splitted to splitledline creating N columns based on splitted length
		splittedLine = append(splittedLine, splitted)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return splittedLine

}

func ReadFromDelimiter(filePath string, delimiter string, offset int) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	results := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if os.Getenv("DEBUG") == "true" {
			fmt.Printf("line: %s\n", line)
		}
		start := 0
		for {
			index := strings.Index(line[start:], delimiter)
			if index == -1 {
				break
			}
			start += index + len(delimiter)
			end := start + offset
			if end > len(line) {
				end = len(line)
			}
			results = append(results, line[start:end])
			start = end
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}

func ReadFileLineByLine(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if os.Getenv("DEBUG") == "true" {
			fmt.Printf("line: %s\n", line)
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func ReadMultipleTypesOfInputs(filePath string, delimiter string) [][]string {
	lines := ReadFileLineByLine(filePath)
	var inputs [][]string
	var currentInput []string

	for _, line := range lines {
		if line == delimiter {
			if len(currentInput) > 0 {
				inputs = append(inputs, currentInput)
				currentInput = []string{}
			}
		} else {
			currentInput = append(currentInput, line)
		}
	}

	if len(currentInput) > 0 {
		inputs = append(inputs, currentInput)
	}

	return inputs
}
