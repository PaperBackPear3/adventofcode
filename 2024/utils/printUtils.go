package utils

import (
	"adventofcode-2024/models"
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func PrintMapWithColors(lines [][]string, locationsToColor map[string][]models.Coordinates) {

	colors := map[string]string{}
	for location := range locationsToColor {
		colors[location] = GenerateRandomColor(colors)
	}

	for _, line := range lines {
		for x := range line {
			char := line[x]
			if char != "." {
				fmt.Printf("\033[38;2;%sm%s\033[0m", colors[char], char)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func GenerateRandomColor(existingColors map[string]string) string {
	var color string
	for {
		rand.Seed(uint64(time.Now().UnixNano()))
		r := rand.Intn(256)
		g := rand.Intn(256)
		b := rand.Intn(256)
		color = fmt.Sprintf("%d;%d;%d", r, g, b)
		if _, exists := existingColors[color]; !exists {
			break
		}
	}
	return color
}
