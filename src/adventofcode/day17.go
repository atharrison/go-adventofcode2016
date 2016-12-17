package adventofcode

import (
	"crypto/md5"
	"fmt"
	"sort"
)

func Day17() {

	//input := "ihgpwlah"
	input := "qtetzkpl"
	bestPath := ""
	aBestPathFound := false

	toTry := Paths{Path{Value: input, xLoc: 0, yLoc: 0}}

	for len(toTry) > 0 {
		next := toTry[0]
		toTry = toTry[1:]

		fmt.Printf("(%v) Checking %v at %v,%v\n", len(toTry), next.Value, next.xLoc, next.yLoc)
		if next.xLoc == 3 && next.yLoc == 3 {
			if len(next.Value) < len(bestPath) || !aBestPathFound {
				fmt.Printf("(%v) Found potential best %v of length %v", len(toTry), next.Value, len(next.Value))
				aBestPathFound = true
				bestPath = next.Value
			}
		}

		openDoors := openDoorsForInput(next.Value)
		for _, d := range openDoors {
			switch d {
			case 'U':
				if next.yLoc > 1 {
					newPath := Path{
						Value: next.Value + "U",
						xLoc:  next.xLoc,
						yLoc:  next.yLoc - 1,
					}
					toTry = append(toTry, newPath)
				}
			case 'D':
				if next.yLoc < 3 {
					newPath := Path{
						Value: next.Value + "D",
						xLoc:  next.xLoc,
						yLoc:  next.yLoc + 1,
					}
					toTry = append(toTry, newPath)
				}
			case 'L':
				if next.xLoc > 1 {
					newPath := Path{
						Value: next.Value + "L",
						xLoc:  next.xLoc - 1,
						yLoc:  next.yLoc,
					}
					toTry = append(toTry, newPath)
				}
			case 'R':
				if next.xLoc < 3 {
					newPath := Path{
						Value: next.Value + "R",
						xLoc:  next.xLoc + 1,
						yLoc:  next.yLoc,
					}
					toTry = append(toTry, newPath)
				}
			}

		}
		sort.Sort(toTry)
	}

	fmt.Printf("Best Path: %v\n", bestPath)
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

	result := []rune{}

	for i := 0; i < 4; i++ {
		toCheck := roomHash[i]
		switch toCheck {
		case 'b', 'c', 'd', 'e', 'f':
			result = append(result, doors[i])
		}
	}
	return result

}
