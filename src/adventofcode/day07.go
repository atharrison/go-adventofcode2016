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

	//Part 2
	var outsideABA []string
	var insideABA []string

	for i := 0; i < len(line); i++ {
		segment += string(line[i])
		if !inside {
			if line[i] == '[' {
				hasabba = hasabba || hasABBA(segment)
				newResults := findABA(segment)
				for _, r := range newResults {
					outsideABA = append(outsideABA, r)
				}

				inside = true
				segment = ""
			}

		} else {
			if line[i] == ']' {
				inside = false
				hasABBAInside = hasABBAInside || hasABBA(segment)
				newResults := findABA(segment)
				for _, r := range newResults {
					insideABA = append(insideABA, r)
				}

				segment = ""
			}
		}
	}
	if !inside {
		hasabba = hasabba || hasABBA(segment)
		newResults := findABA(segment)
		for _, r := range newResults {
			outsideABA = append(outsideABA, r)
		}
	} else {
		hasABBAInside = hasABBAInside || hasABBA(segment)
		newResults := findABA(segment)
		for _, r := range newResults {
			insideABA = append(insideABA, r)
		}
	}

	// Part 1
	//if hasabba && !hasABBAInside {
	//	fmt.Printf("TLS: %v\n", line)
	//	return 1
	//}
	//return 0

	// Part 2
	if hasSSL(outsideABA, insideABA) {
		return 1
	}
	return 0

}

func hasSSL(outside []string, inside []string) bool {
	fmt.Printf("Checking SSL %v against %v\n", outside, inside)
	for _, item := range outside {
		toCheck := string(item[1]) + string(item[0]) + string(item[1])
		for _, insideItem := range inside {
			if toCheck == insideItem {
				return true
			}
		}
	}
	return false
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

func findABA(segment string) []string {
	var results []string

	fmt.Printf("Checking %v\n", segment)
	for i := 0; i < len(segment)-2; i++ {
		if segment[i] == segment[i+2] && segment[i] != segment[i+1] && segment[i+2] != '[' && segment[i+2] != ']' {
			fmt.Printf("ABA: %v\n", segment)
			results = append(results, segment[i:i+3])
		}

	}
	return results
}
