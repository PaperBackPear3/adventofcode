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
			if !isCloseShade(existingColors, color) {
				break
			}
			break
		}
	}
	return color
}

func isCloseShade(existingColors map[string]string, color string) bool {
	const threshold = 200

	r, g, b := parseColor(color)
	newColor := [3]int{r, g, b}
	for _, existingColor := range existingColors {
		r, g, b := parseColor(existingColor)
		existing := [3]int{r, g, b}
		if colorDistance(newColor, existing) < threshold {
			return true
		}
	}
	return false
}

func parseColor(color string) (int, int, int) {
	var r, g, b int
	fmt.Sscanf(color, "%d;%d;%d", &r, &g, &b)
	return r, g, b
}

func colorDistance(c1, c2 [3]int) int {
	rDiff := c1[0] - c2[0]
	gDiff := c1[1] - c2[1]
	bDiff := c1[2] - c2[2]
	return rDiff*rDiff + gDiff*gDiff + bDiff*bDiff
}
