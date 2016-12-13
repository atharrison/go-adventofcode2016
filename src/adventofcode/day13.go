package adventofcode

import (
	"crypto/md5"
	"fmt"
	"sort"
)

const maxSize = 75

//const maxSize = 10
//const favoriteNumber = 10
const favoriteNumber = 1352

var traveled = make(map[string]*Location)

func Day13() {

	//goalX, goalY := 7, 4
	goalX, goalY := 31, 39

	floorplan := make([][]string, maxSize)
	for y := 0; y < maxSize; y++ {
		floorplan[y] = make([]string, maxSize)
		for x := 0; x < maxSize; x++ {
			val := Day13Formula(x, y)

			binaryArray := ConvertToBase(int64(val), 2)
			hasOdd := ArrayHasOdd(binaryArray)
			fmt.Printf("%v is %v, hasOdd? %v\n", val, binaryArray, hasOdd)
			if hasOdd {
				floorplan[y][x] = "#"
			} else {
				floorplan[y][x] = "."
			}

		}
	}

	//floorplan[goalY][goalX] = "G"
	PrintFloorplan(floorplan)

	fmt.Println(WalkFloor(floorplan, goalX, goalY))

	// Part 2 (I was already tracking 'traveled', so just needed to ask what I tracked
	shorts := 0
	for _, v := range traveled {
		if v.Dist <= 50 {
			shorts += 1
		}
	}
	fmt.Printf("50 or better: %v\n", shorts)
}

func PrintFloorplan(floorplan [][]string) {
	for y := 0; y < maxSize; y++ {
		for x := 0; x < maxSize; x++ {
			fmt.Printf(floorplan[y][x])
		}
		fmt.Println()
	}
}

func ArrayHasOdd(arr []int) bool {
	total := 0
	for _, v := range arr {
		if v%2 == 1 {
			total += 1
		}
	}
	return (total % 2) == 1
}

func Day13Formula(x int, y int) int {
	offset := favoriteNumber
	return x*x + 3*x + 2*x*y + y + y*y + offset
}

type Location struct {
	X    int
	Y    int
	Dist int
}

func (loc *Location) Hash() string {
	toHash := fmt.Sprintf("%v:%v", loc.X, loc.Y)
	hash := fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
	return hash
}

type Locations []*Location

func (slice Locations) Len() int {
	return len(slice)
}

func (slice Locations) Less(i, j int) bool {
	return slice[i].Dist < slice[j].Dist
}

func (slice Locations) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func WalkFloor(floorplan [][]string, goalX, goalY int) int {

	lowestDist := 9999
	stepsToProcess := Locations{
		&Location{
			X:    1,
			Y:    1,
			Dist: 0,
		},
	}

	for len(stepsToProcess) > 0 {

		next := stepsToProcess[0]
		stepsToProcess = stepsToProcess[1:]
		fmt.Printf("Best: %v\t ToProcess: %v\tNow at [%v, %v] D:%v    \n", lowestDist, len(stepsToProcess), next.X, next.Y, next.Dist)
		if next.Dist > lowestDist {
			continue
		}

		if next.X == goalX && next.Y == goalY && lowestDist > next.Dist {
			lowestDist = next.Dist
			fmt.Printf("Best Dist now %v\n", lowestDist)
		}

		for xDelta := -1; xDelta < 2; xDelta++ {
			if next.X+xDelta < 0 || next.X+xDelta >= maxSize {
				continue
			}
			for yDelta := -1; yDelta < 2; yDelta++ {
				if next.Y+yDelta < 0 || next.Y+yDelta >= maxSize {
					continue
				}

				if (xDelta == 0 && yDelta == 0) ||
					(xDelta == -1 && yDelta == -1) ||
					(xDelta == -1 && yDelta == 1) ||
					(xDelta == 1 && yDelta == -1) ||
					(xDelta == 1 && yDelta == 1) {
					continue // No diagonals
				}

				if floorplan[next.Y+yDelta][next.X+xDelta] == "." {

					newLoc := &Location{
						X:    next.X + xDelta,
						Y:    next.Y + yDelta,
						Dist: next.Dist + 1,
					}

					h := newLoc.Hash()
					if loc, ok := traveled[h]; ok {
						if loc.Dist > newLoc.Dist {
							traveled[h] = newLoc
							stepsToProcess = append(stepsToProcess, newLoc)
						}
					} else {
						stepsToProcess = append(stepsToProcess, newLoc)
						traveled[h] = newLoc
					}

				}
			}
		}
		sort.Sort(stepsToProcess)
	}

	return lowestDist
}
