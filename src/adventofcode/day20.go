package adventofcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day20() {
	day := "20"
	filename := fmt.Sprintf("data/day%vinput.txt", day)
	input := readFileAsLines(filename)

	ranges := BuildIPRange(input)

	PrintRanges(ranges)

	for i := 0; i < len(ranges)-1; i++ {
		if (ranges[i].end + 1) < ranges[i+1].start {
			//fmt.Printf("%v -> %v\n", ranges[i], ranges[i+1])
			fmt.Printf("Range %v leaves gap before Range %v\n", ranges[i], ranges[i+1])
			fmt.Printf("Found it: %v\n", ranges[i].end+1)
			break
		}
	}
	fmt.Println("Done 20")

	//naive:
	//values := map[int]bool{}
	//for i := 0; i < maxVal; i++ {
	//	values[i] = true
	//}
	//
	//for _, line := range input {
	//	tokens := strings.Split(line, "-")
	//	fmt.Printf("Line: %v\n", line)
	//
	//	smallVal, _ := strconv.Atoi(tokens[0])
	//	largeVal, _ := strconv.Atoi(tokens[1])
	//	for i := smallVal; i < largeVal+1; i++ {
	//		values[i] = false
	//	}
	//
	//}
	//
	//var idx int
	//for idx =0; !values[idx]; idx++ {
	//
	//}
	//
	//fmt.Printf("%v= %v\n", idx, values[idx])
}

type Range struct {
	start int
	end   int
}

func (r *Range) String() string {
	return fmt.Sprintf("%v-%v", r.start, r.end)
}

func (r *Range) Overlaps(other *Range) bool {
	if r.end < other.start || other.end < r.start {
		return false
	}
	return true
}

func (r *Range) Contains(other *Range) bool {
	return r.start >= other.start && r.end >= other.end
}

func CombineRanges(r1 *Range, r2 *Range) *Range {

	newStart := int(math.Min(float64(r1.start), float64(r2.start)))
	newEnd := int(math.Max(float64(r1.end), float64(r2.end)))
	return &Range{
		start: newStart,
		end:   newEnd,
	}
}

func NewRange(line string) *Range {
	tokens := strings.Split(line, "-")
	smallVal, _ := strconv.Atoi(tokens[0])
	largeVal, _ := strconv.Atoi(tokens[1])
	return &Range{
		start: smallVal,
		end:   largeVal,
	}
}

func BuildIPRange(input []string) []*Range {
	var ranges []*Range
	for _, line := range input {
		r := NewRange(line)
		ranges = InsertRange(r, ranges)
		//copy(ranges, InsertRange(r, ranges))
	}
	return ranges
}

func PrintRanges(ranges []*Range) {
	for i := 0; i < len(ranges); i++ {
		fmt.Printf("%v\n", ranges[i])
	}
}

func InsertRange(r *Range, ranges []*Range) []*Range {

	if len(ranges) == 0 {
		fmt.Printf("(%v)\tInserting %v first\n", len(ranges), r)
		return []*Range{r}
	}

	var newRanges []*Range
	appended := false

	if r.start == 0 {
		fmt.Printf("(%v)\tInserting %v at front\n", len(ranges), r)
		if r.Overlaps(ranges[0]) {
			fmt.Printf("(%v)\tCombining %v with %v\n", len(ranges), r, ranges[0])
			newRange := CombineRanges(r, ranges[0])
			newRanges = append(newRanges, newRange)
			newRanges = append(newRanges, ranges[1:]...)
		} else {
			newRanges = append(newRanges, r)
			newRanges = append(newRanges, ranges[0:]...)
		}
		fmt.Printf("(%v)\tAfter Front insert, %v then %v\n", len(ranges), ranges[0], ranges[1])
		return newRanges
	} //short-circuit

	for i := 0; i < len(ranges); i++ {
		if r.Overlaps(ranges[i]) {
			//PrintRanges(ranges)
			//os.Exit(1)
			fmt.Printf("(%v)\tCombining %v with %v\n", len(ranges), r, ranges[i])
			newRange := CombineRanges(r, ranges[i])
			if i == 0 {
				fmt.Printf("(%v)\tInserting Combined %v at front\n", len(ranges), newRange)
				newRanges = []*Range{newRange}
				newRanges = append(newRanges, ranges[1:]...)
				//fmt.Printf("(%v)\tAfter Combined insert\n", len(newRanges))
				appended = true
				break
			} else if i == len(ranges)-1 {
				fmt.Printf("(%v)\tInserting Combined %v at end\n", len(ranges), newRange)
				newRanges = ranges[0 : i-1]
				newRanges = append(newRanges, newRange)
				//fmt.Printf("(%v)\tAfter Combined insert\n", len(newRanges))
				appended = true
				break
			} else {
				fmt.Printf("(%v)\tInserting Combined %v before %v\n", len(ranges), newRange, ranges[i+1])
				newRanges = ranges[0:i]
				newRanges = append(newRanges, newRange)
				newRanges = append(newRanges, ranges[i+1:]...)
				//fmt.Printf("(%v)\tAfter Combined insert\n", len(newRanges))
				appended = true
				break
			}
		} else if r.end < ranges[i].start {
			if i == 0 {
				fmt.Printf("(%v)\tInserting %v at start, in front of %v\n", len(ranges), r, ranges[i])
				newRanges = []*Range{r}
				newRanges = append(newRanges, ranges...)
				appended = true
				//PrintRanges(newRanges)
				//os.Exit(1)
				break
			} else {
				fmt.Printf("(%v)\tInserting %v before %v\n", len(ranges), r, ranges[i])
				for _, rr := range ranges[0:i] {
					newRanges = append(newRanges, rr)
				}
				//newRanges = ranges[0:i]
				newRanges = append(newRanges, r)
				newRanges = append(newRanges, ranges[i:]...)
				//PrintRanges(newRanges)
				//os.Exit(1)
				appended = true
				break
			}
		}
	}

	if !appended {
		fmt.Printf("(%v)\tInserting %v at end\n", len(ranges), r)
		for _, rr := range ranges {
			newRanges = append(newRanges, rr)
		}
		newRanges = append(newRanges, r)
	}
	return newRanges
}
