package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
