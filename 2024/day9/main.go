package main

import (
	"adventofcode-2024/utils"
)

type Values struct {
	fileId int
	length int
}

type Blocks struct {
	values      Values
	spaceLength int
}

func main() {
	line := utils.ReadFileLineByLine("./input.txt")[0]
	blocks := []Blocks{}
	counter := 0
	for index := 0; index < len(line); index++ {
		val := Values{}
		if index%2 == 0 {
			val = Values{fileId: counter, length: utils.Atoi(string(line[index]))}
			blocks = append(blocks, Blocks{values: val})
			counter = counter + 1
		} else {
			blocks[len(blocks)-1].spaceLength = utils.Atoi(string(line[index]))
		}

	}
	checkSumIds := []int{}
	for i := 0; i < len(blocks); i++ {
		blocksLen := blocks[i].values.length
		for loops := 0; loops < blocksLen; loops++ {
			checkSumIds = append(checkSumIds, blocks[i].values.fileId)
			blocks[i].values.length = blocks[i].values.length - 1
		}
		//todo remember from where to start, should not be the end each time
		for fromEnd := len(blocks) - 1; fromEnd >= 0; fromEnd-- {
			if blocks[i].spaceLength == 0 {
				break
			}
			// fromEndBlock := blocks[fromEnd]
			if blocks[fromEnd].values.length > 0 {
				for loops := 0; loops <= blocks[fromEnd].values.length+1; loops++ {
					if blocks[i].spaceLength > 0 {
						checkSumIds = append(checkSumIds, blocks[fromEnd].values.fileId)
						blocks[fromEnd].values.length = blocks[fromEnd].values.length - 1
						blocks[i].spaceLength = blocks[i].spaceLength - 1
					} else {
						break
					}
				}
			}
		}
	}
	println(checkSumIds)

	sum := int64(0)
	for i, val := range checkSumIds {
		sum += int64(val * i)
	}

	println(blocks)

}
