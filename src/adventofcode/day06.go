package adventofcode

import (
	"fmt"
	"strings"
)

func Day06() {
	day := "06"
	filename := fmt.Sprintf("data/day%vinput.txt", day)
	input := readFileAsLines(filename)

	line1 := input[0]
	colArrayMap := make([]map[string]int, len(line1))
	for i := 0; i < len(line1); i++ {
		colArrayMap[i] = map[string]int{}
	}

	for _, line := range input {
		countCharsPerColumn(line, colArrayMap)
	}

	for i := 0; i < len(line1); i++ {
		//letter := maxOfDay6(colArrayMap[i]) // Part 1
		letter := leastOfDay6(colArrayMap[i]) // Part 2
		fmt.Printf(letter)
	}
}

func maxOfDay6(col map[string]int) string {
	total := 0
	letter := ""
	for k, v := range col {
		if v > total {
			letter = k
			total = v
		}
	}
	return letter
}

func leastOfDay6(col map[string]int) string {
	total := 999
	letter := ""
	for k, v := range col {
		if v < total {
			letter = k
			total = v
		}
	}
	return letter
}

func countCharsPerColumn(line string, colArrayMap []map[string]int) {
	for i, c := range strings.TrimSpace(line) {
		cStr := string(c)
		col := colArrayMap[i]
		col[cStr] += 1
	}
}
