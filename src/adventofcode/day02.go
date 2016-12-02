package adventofcode

import (
	"fmt"
	"strings"
)

func Day02() {
	fmt.Println("Day01 START")

	input := readFileAsLines("data/day02input.txt")

	code := []int{}
	for _, line := range input {
		fmt.Printf("Processing %v\n", line)
		code = append(code, processElevatorLine(strings.TrimSpace(line)))
	}
	fmt.Printf("Code: %v", code)
}

var x = 1
var y = 1

var keypad = [][]int{
	[]int{7, 4, 1},
	[]int{8, 5, 2},
	[]int{9, 6, 3},
}

func processElevatorLine(line string) int {
	tokens := strings.Split(line, "")
	fmt.Printf("Tokens: %v\n", tokens)
	for _, val := range tokens {
		switch strings.Trim(val, "\n") {
		case "R":
			if x < 2 {
				fmt.Printf("Moving R from %v to %v\n", keypad[x][y], keypad[x+1][y])
				x += 1
			} else {
				fmt.Printf("Skipping R\n")
			}
		case "L":
			if x > 0 {
				fmt.Printf("Moving L from %v to %v\n", keypad[x][y], keypad[x-1][y])
				x -= 1
			} else {
				fmt.Printf("Skipping L\n")
			}
		case "U":
			if y < 2 {
				fmt.Printf("Moving U from %v to %v\n", keypad[x][y], keypad[x][y+1])
				y += 1
			} else {
				fmt.Printf("Skipping U\n")
			}
		case "D":
			if y > 0 {
				fmt.Printf("Moving D from %v to %v\n", keypad[x][y], keypad[x][y-1])
				y -= 1
			} else {
				fmt.Printf("Skipping D\n")
			}
		}
	}
	return keypad[x][y]
}
