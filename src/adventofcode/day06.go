package adventofcode

import (
	"fmt"
	"strings"
)

func Day06() {
	day := "06"
	filename := fmt.Sprintf("data/day%vinput.txt", day)
	input := readFileAsLines(filename)

	//var colArrayMap []map[string]int

	line1 := input[0]
	colArrayMap := []map[string]int{
		map[string]int{},
		map[string]int{},
		map[string]int{},
		map[string]int{},
		map[string]int{},
		map[string]int{},
		map[string]int{},
		map[string]int{},
	}

	for i := 0; i < len(line1); i++ {
		colArrayMap[i] = map[string]int{}
	}

	fmt.Printf("ColArrayMap: %v\n", colArrayMap)
	//for i := 0; i < len(line1); i++ {
	//	colArrayMap[i] = make(map[string]int)
	//}
	for _, line := range input {
		//fmt.Printf("Line: %v\n", line)
		countCharsPerColumn(line, colArrayMap)
	}

	for i := 0; i < len(line1); i++ {
		letter := maxOf(colArrayMap[i])
		fmt.Printf(letter)
	}

	//fmt.Printf("%v\n", colArrayMap)
}

func maxOf(col map[string]int) string {
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

func countCharsPerColumn(line string, colArrayMap []map[string]int) {
	fmt.Printf("Processing [%v]\n", line)
	for i, c := range strings.TrimSpace(line) {

		cStr := string(c)
		//if _, ok := colArrayMap[i]; !ok {
		//	colArrayMap[i] = map[string]int{}
		//}
		fmt.Printf("Index %v, char %v\n", i, cStr)
		fmt.Printf("Getting map %v: %v\n", i, colArrayMap[i])
		col := colArrayMap[i]
		col[cStr] += 1
		//if _, ok := col[cStr]; !ok {
		//	col[cStr] = 1
		//} else {
		//	col[cStr]+= 1
		//}
	}
}
