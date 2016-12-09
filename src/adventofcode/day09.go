package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

const SCAN = 1
const DECOMP = 2

var day09Input string

func Day09() {
	day := "09"
	filename := fmt.Sprintf("data/day%vinput.txt", day)
	day09Input = strings.TrimSpace(readFileAsString(filename))

	//fmt.Printf("%v\n", day09Input)
	result := decompressInputWithPointers(0, len(day09Input), 0)
	//resultCount := len(result)
	//var result string

	//fmt.Println(result)
	//fmt.Println(len(result))
	fmt.Printf("\n%v\n", result)
}

func decompressInputWithPointers(ptr int, end int, count int64) int64 {

	//fmt.Printf("Decompressing from %v, %v...\n", ptr, day09Input[ptr:ptr+3])

	state := SCAN
	var decompInstr string
	var endPtr int
	for ptr < end {
		fmt.Printf("Count: %v\tPtr: %v                          \r", count, ptr)
		switch state {
		case SCAN:
			if day09Input[ptr] == '(' {
				state = DECOMP
				endPtr = ptr
				found := false
				decompInstr = ""
				for !found {
					endPtr += 1
					if day09Input[endPtr] == ')' {
						found = true
						//fmt.Printf("Building Instr: %v\n", decompInstr)
						//decompInstr += string(day09Input[endPtr])
					} else {
						//fmt.Printf("Building Instr: %v\n", decompInstr)
						decompInstr += string(day09Input[endPtr])
					}
				}
			} else {
				//fmt.Printf("Counting %v once\n", string(day09Input[ptr]))
				count += 1
			}
			ptr += 1
		case DECOMP:

			//fmt.Printf("New Decompress directive: %v\n", decompInstr)
			tokens := strings.Split(decompInstr, "x")
			length, _ := strconv.Atoi(tokens[0])
			repeat, _ := strconv.Atoi(tokens[1])

			sectionCount := decompressSectionCount(endPtr+1, endPtr+length+1, length, repeat, count)
			//fmt.Printf("Section Count %v added to count %v\n", sectionCount, count)
			count += sectionCount
			ptr = endPtr + length + 1
			state = SCAN
		}
	}

	//fmt.Printf("Decompress count now %v\n", count)
	return count
}

func decompressSectionCount(start int, end int, length, repeat int, count int64) int64 {
	var result int64
	//fmt.Printf("Decompressing Section %v-%v, next %v chars %v times (count: %v)...\n", start, end, length, repeat, count)

	nextDecomprStr := day09Input[start:end]
	//fmt.Printf("NextDecompressStr: %v\n", nextDecomprStr)
	if strings.Index(nextDecomprStr, "(") == -1 {
		//fmt.Printf("Adding %v*%v to %v\n", length, repeat, count)
		result = int64(length * repeat)
	} else {
		for j := 0; j < repeat; j++ {
			result = decompressInputWithPointers(start, end, result)
		}
		//result = count
	}

	//fmt.Printf("Section count now %v\n", result)
	return result
}

// Below here be dragons. It will find the solution, but you will need way more memory than you have for Part 2.
//==============================================================================================================
//func decompressInputBrute(input string) string {
//	fmt.Printf("\nDecompressing input of length %v\n", len(input))
//	var result string
//	//var resultCount int64
//	var state = SCAN
//	var decompInstr string
//	for i := 0; i < len(input); i++ {
//		next := input[i]
//		switch state {
//		case SCAN:
//			if next == '(' {
//				state = DECOMP
//				decompInstr = ""
//				found := false
//
//				for !found {
//					i += 1
//					if input[i] == ')' {
//						found = true
//					} else {
//						decompInstr += string(input[i])
//					}
//				}
//			} else {
//				result += string(next)
//				//resultCount += 1
//			}
//		case DECOMP:
//			tokens := strings.Split(decompInstr, "x")
//			length, _ := strconv.Atoi(tokens[0])
//			repeat, _ := strconv.Atoi(tokens[1])
//
//			result += decompressSection(input[i:i+length], length, repeat)
//			//resultCount += len(result)
//			i = i + length - 1
//			state = SCAN
//
//		}
//	}
//	return result
//}
//
//func decompressSection(section string, length, repeat int) string {
//	var result string
//	fmt.Printf("Decompressing next %v chars %v times...\r", length, repeat)
//	for r := 0; r < repeat; r++ {
//		result += section
//		//resultCount += int64(length)
//	}
//
//
//	if strings.Index(result, "(") == -1 {
//		return result
//	} else {
//		result = decompressInputBrute(result)
//	}
//
//	return result
//}
