package adventofcode

import "fmt"

func Day00() {
	day := "00"
	filename := fmt.Sprintf("data/day%vinput.txt", day)
	input := readFileAsLines(filename)

	for _, line := range input {
		fmt.Printf("Line: %v\n", line)
	}

	fmt.Printf("%v\n", input)
}
