package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

func Day03() {
	fmt.Println("Day03 START")

	input := readFileAsLines("data/day03input.txt")

	//fmt.Printf("Lines: %v", input)

	total := 0
	for _, line := range input {
		total += isPossibleTriangle(line)
	}
	fmt.Printf("Total: %v", total)
}

func isPossibleTriangle(line string) int {

	tokens := strings.Split(line, " ")

	largest := 0
	sum := 0

	fmt.Printf("line: %v\nTokens: %v", line, tokens)
	for _, token := range tokens {
		val, _ := strconv.ParseInt(strings.TrimSpace(token), 10, 64)
		if val == 0 {
			continue
		}
		fmt.Printf("%v ", val)
		if int(val) > largest {
			largest = int(val)
		}
		sum += int(val)
	}
	fmt.Printf("|\nLargest greater than other? %v, %v\t%v\n", largest, sum-largest, largest > (sum-largest))
	if largest >= (sum - largest) {
		return 0
	}
	fmt.Printf("Possible: %v\n", tokens)
	return 1
}
