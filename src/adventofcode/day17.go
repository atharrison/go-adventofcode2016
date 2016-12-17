package adventofcode

import (
	"crypto/md5"
	"fmt"
	"sort"
)

func Day17() {

	//input := "ihgpwlah"
	//input := "kglvqrro"
	//input := "ulqzkmiv"
	input := "qtetzkpl"
	bestPath := ""
	aBestPathFound := false

	// Part 2
	longestPath := ""

	toTry := Paths{Path{Value: input, xLoc: 0, yLoc: 0}}

	for len(toTry) > 0 {
		next := toTry[0]
		toTry = toTry[1:]

		fmt.Printf("(%v) Checking %v at %v,%v\n", len(toTry), next.Value, next.xLoc, next.yLoc)
		if next.xLoc == 3 && next.yLoc == 3 {
			if len(next.Value) < len(bestPath) || !aBestPathFound {
				fmt.Printf("(%v) Found potential best %v of length %v\n", len(toTry), next.Value, len(next.Value))
				bestPath = next.Value
			}
			if !aBestPathFound {
				aBestPathFound = true
			}

			// Part 2:
			if len(next.Value) > len(longestPath) {
				longestPath = next.Value
			}
			continue
		}

		openDoors := openDoorsForInput(next.Value)

		for _, d := range openDoors {
			//fmt.Printf("For %v, adding %v\n", next.Value, string(d))
			switch d {
			case 'U':
				if next.yLoc > 0 {
					newPath := Path{
						Value: next.Value + "U",
						xLoc:  next.xLoc,
						yLoc:  next.yLoc - 1,
					}
					toTry = append(toTry, newPath)
					//fmt.Printf("Adding U to Paths, now len (%v)\n", len(toTry))
				}
			case 'D':
				if next.yLoc < 3 {
					newPath := Path{
						Value: next.Value + "D",
						xLoc:  next.xLoc,
						yLoc:  next.yLoc + 1,
					}
					toTry = append(toTry, newPath)
					//fmt.Printf("Adding D to Paths, now len (%v)\n", len(toTry))
				}
			case 'L':
				if next.xLoc > 0 {
					newPath := Path{
						Value: next.Value + "L",
						xLoc:  next.xLoc - 1,
						yLoc:  next.yLoc,
					}
					toTry = append(toTry, newPath)
					//fmt.Printf("Adding L to Paths, now len (%v)\n", len(toTry))
				}
			case 'R':
				if next.xLoc < 3 {
					newPath := Path{
						Value: next.Value + "R",
						xLoc:  next.xLoc + 1,
						yLoc:  next.yLoc,
					}
					toTry = append(toTry, newPath)
					//fmt.Printf("Adding R to Paths, now len (%v)\n", len(toTry))
				}
			}

		}
		//fmt.Printf("toTry now len (%v)\n", len(toTry))
		sort.Sort(toTry)
	}

	fmt.Printf("Best Path: %v\n", bestPath)
	fmt.Printf("Len Longest Path: %v\n", len(longestPath)-len(input))
}

type Paths []Path

type Path struct {
	Value string
	xLoc  int
	yLoc  int
}

func (p Paths) Len() int { return len(p) }
func (p Paths) Less(i, j int) bool {
	return len(p[i].Value) < len(p[j].Value)
}
func (p Paths) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

var doors = []rune{'U', 'D', 'L', 'R'}

func openDoorsForInput(input string) []rune {
	roomHash := fmt.Sprintf("%x", md5.Sum([]byte(input)))
	//fmt.Printf("hash: %v", roomHash[0:4])

	result := []rune{}

	for i := 0; i < 4; i++ {
		toCheck := roomHash[i]
		switch toCheck {
		case 'b', 'c', 'd', 'e', 'f':
			result = append(result, doors[i])
		}
	}
	//fmt.Printf(" adds %v\n", result)
	return result

}
