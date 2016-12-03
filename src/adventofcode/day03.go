package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

func Day03() {
	fmt.Println("Day03 START")

	// Part 1
	//input := readFileAsLines("data/day03input.txt")

	//fmt.Printf("Lines: %v", input)

	//total := 0
	//for _, line := range input {
	//	total += isPossibleTriangle(line)
	//}
	//fmt.Printf("Total: %v", total)

	// Part 2
	input := readFileAsLines("data/day03input.txt")
	tokens := splitTriangleInputToTokens(input)

	total := findVerticalTriangles(tokens)
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

//Part 2

func splitTriangleInputToTokens(lines []string) []int {
	var values []int
	for _, line := range lines {
		t1 := line[2:5]
		t2 := line[7:10]
		t3 := line[12:15]
		fmt.Printf("%v|%v|%v\n", t1, t2, t3)
		v1, _ := strconv.ParseInt(strings.TrimSpace(t1), 10, 64)
		v2, _ := strconv.ParseInt(strings.TrimSpace(t2), 10, 64)
		v3, _ := strconv.ParseInt(strings.TrimSpace(t3), 10, 64)
		values = append(values, int(v1))
		values = append(values, int(v2))
		values = append(values, int(v3))
	}

	return values
}

func findVerticalTriangles(values []int) int {
	total := 0
	fmt.Println(values)
	for i := 0; i < len(values)-8; i += 9 {
		total += isTriangle([]int{values[i], values[i+3], values[i+6]})
		total += isTriangle([]int{values[i+1], values[i+4], values[i+7]})
		total += isTriangle([]int{values[i+2], values[i+5], values[i+8]})
	}
	return total
}

func isTriangle(sides []int) int {

	fmt.Printf("Checking %v\n", sides)
	largest := 0
	sum := 0
	for _, s := range sides {

		if s > largest {
			largest = s
		}
		sum += s
	}
	if largest >= (sum - largest) {
		return 0
	}
	return 1

}
