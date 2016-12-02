package adventofcode

import (
	"fmt"
	"strings"
)

func Day02() {
	fmt.Println("Day02 START")

	input := readFileAsLines("data/day02input.txt")

	code := []string{}
	for _, line := range input {
		fmt.Printf("Processing %v\n", line)
		code = append(code, processElevatorLine(strings.TrimSpace(line)))
	}
	fmt.Printf("Code: %v", code)
}

// Part 1
//var x = 1
//var y = 1

// Part 1 Keypad
//var keypad = [][]int{
//	[]int{7, 4, 1},
//	[]int{8, 5, 2},
//	[]int{9, 6, 3},
//}

// Part 2
var x = 0
var y = 2

//Part 2 Keypad
/*
    1
  2 3 4
5 6 7 8 9
  A B C
    D
*/
var keypad = [][]rune{
	[]rune{'x', 'x', '5', 'x', 'x'},
	[]rune{'x', 'A', '6', '2', 'x'},
	[]rune{'D', 'B', '7', '3', '1'},
	[]rune{'x', 'C', '8', '4', 'x'},
	[]rune{'x', 'x', '9', 'x', 'x'},
}

func processElevatorLine(line string) string {
	tokens := strings.Split(line, "")
	fmt.Printf("Tokens: %v\n", tokens)
	for _, val := range tokens {
		switch strings.Trim(val, "\n") {
		case "R":
			//if x < 2 {
			if x < 4 && keypad[x+1][y] != 'x' {
				fmt.Printf("Moving R from %v to %v\n", string(keypad[x][y]), string(keypad[x+1][y]))
				x += 1
			} else {
				fmt.Printf("Skipping R\n")
			}
		case "L":
			//if x > 0 {
			if x > 0 && keypad[x-1][y] != 'x' {
				fmt.Printf("Moving L from %v to %v\n", string(keypad[x][y]), string(keypad[x-1][y]))
				x -= 1
			} else {
				fmt.Printf("Skipping L\n")
			}
		case "U":
			//if y < 2 {
			if y < 4 && keypad[x][y+1] != 'x' {
				fmt.Printf("Moving U from %v to %v\n", string(keypad[x][y]), string(keypad[x][y+1]))
				y += 1
			} else {
				fmt.Printf("Skipping U\n")
			}
		case "D":
			//if y > 0 {
			if y > 0 && keypad[x][y-1] != 'x' {
				fmt.Printf("Moving D from %v to %v\n", string(keypad[x][y]), string(keypad[x][y-1]))
				y -= 1
			} else {
				fmt.Printf("Skipping D\n")
			}
		}
	}
	return string(keypad[x][y])
}
