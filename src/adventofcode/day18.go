package adventofcode

import (
	"fmt"
	"strings"
)

func Day18() {

	input := ".^^.^^^..^.^..^.^^.^^^^.^^.^^...^..^...^^^..^^...^..^^^^^^..^.^^^..^.^^^^.^^^.^...^^^.^^.^^^.^.^^.^."
	//input := ".^^.^.^^^^"
	//maxRows := 40
	//maxRows := 10

	// Part 2: (efficient enough, just don't save all prev rows, and don't output until finished)
	maxRows := 400000

	row := strings.Split(input, "")
	//fmt.Printf("%v\n", row)

	//rows := make([][]string, maxRows)
	//rows[0] = row
	nextRow := row
	totalSafe := countSafe(row)
	//fmt.Printf("%v (%v)\n", row, totalSafe)
	for rowIdx := 1; rowIdx < maxRows; rowIdx++ {
		//prevRow := rows[rowIdx-1]
		prevRow := nextRow

		nextRow = make([]string, len(prevRow))
		for x := 0; x < len(prevRow); x++ {
			if IsNextTileSafe(prevRow, x) {
				nextRow[x] = "."
			} else {
				nextRow[x] = "^"
			}
		}
		totalSafe += countSafe(nextRow)
		//rows[rowIdx] = nextRow
		//fmt.Printf("%v (%v)\n", nextRow, totalSafe)
	}
	//fmt.Println(rows[maxRows-1])

	fmt.Printf("Safe: %v\n", totalSafe)
}

func countSafe(row []string) int {
	safe := 0
	for _, v := range row {
		switch v {
		case ".":
			safe += 1
		}

	}
	return safe
}

func IsNextTileSafe(prevRow []string, idx int) bool {
	var left bool
	if idx == 0 {
		left = true
	} else {
		left = TileSafe(prevRow[idx-1])
	}

	var right bool
	if idx == len(prevRow)-1 {
		right = true
	} else {
		right = TileSafe(prevRow[idx+1])
	}
	center := TileSafe(prevRow[idx])

	if !left && !center && right {
		return false
	} else if !center && !right && left {
		return false
	} else if !left && right && center {
		return false
	} else if !right && left && center {
		return false
	}
	return true
}

func TileSafe(tile string) bool {
	switch tile {
	case ".":
		return true
	default:
		return false
	}
}
