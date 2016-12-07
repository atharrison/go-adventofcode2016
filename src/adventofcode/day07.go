package adventofcode

import "fmt"

func Day07() {
	day := "07"
	filename := fmt.Sprintf("data/day%vinput.txt", day)
	input := readFileAsLines(filename)

	total := 0
	for _, line := range input {
		//for i:=0; i< 2; i++ {
		//	line := input[i]
		fmt.Printf("Line: %v\n", line)
		total += isIpV7TLS(line)
	}

	fmt.Printf("Total: %v\n", total)
}

func isIpV7TLS(line string) int {
	inside := false
	segment := ""
	hasabba := false
	hasABBAInside := false
	for i := 0; i < len(line); i++ {
		segment += string(line[i])
		if !inside {
			if line[i] == '[' {
				hasabba = hasabba || hasABBA(segment)
				inside = true
				segment = ""
			}

		} else {
			if line[i] == ']' {
				inside = false
				hasABBAInside = hasABBAInside || hasABBA(segment)
				segment = ""
			}
		}
	}
	if !inside {
		hasabba = hasabba || hasABBA(segment)
		//inside = true
		//segment = ""

	} else {
		//inside = false
		hasABBAInside = hasABBAInside || hasABBA(segment)
		//segment = ""
	}
	if hasabba && !hasABBAInside {
		fmt.Printf("TLS: %v\n", line)
		return 1
	}
	return 0
}

func hasABBA(segment string) bool {
	fmt.Printf("Checking %v\n", segment)
	for i := 0; i < len(segment)-3; i++ {
		if segment[i] == segment[i+3] && segment[i+1] == segment[i+2] && segment[i] != segment[i+1] && segment[i+3] != '[' && segment[i+3] != ']' {
			fmt.Printf("ABBA: %v\n", segment)
			return true
		}

	}
	return false
}
