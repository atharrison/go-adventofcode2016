package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

const SCAN = 1
const DECOMP = 2

func Day09() {
	day := "09"
	filename := fmt.Sprintf("data/day%vinput.txt", day)
	input := strings.TrimSpace(readFileAsString(filename))

	fmt.Printf("%v\n", input)

	var result string
	var state = SCAN
	var decompInstr string
	for i := 0; i < len(input); i++ {
		next := input[i]
		switch state {
		case SCAN:
			if next == '(' {
				state = DECOMP
				decompInstr = ""
				found := false

				for !found {
					i += 1
					if input[i] == ')' {
						found = true
					} else {
						decompInstr += string(input[i])
					}
				}
			} else {
				result += string(next)
			}
		case DECOMP:
			tokens := strings.Split(decompInstr, "x")
			length, _ := strconv.Atoi(tokens[0])
			repeat, _ := strconv.Atoi(tokens[1])

			fmt.Printf("Decompressing next %v chars %v times...\n", length, repeat)
			for r := 0; r < repeat; r++ {
				result += input[i : i+length]
				//for j := 0; j < length; j++ {
				//	result += string(input[i+j])
				//}
			}
			i = i + length - 1
			state = SCAN

		}
	}

	fmt.Println(result)
	fmt.Println(len(result))
}
