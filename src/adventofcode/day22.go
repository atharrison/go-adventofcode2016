package adventofcode

import (
	"crypto/md5"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var maxX int
var maxY int

var ZeroPoint *Point

func Day22() {
	day := "22"
	//filename := fmt.Sprintf("data/day%vinput.txt", day)
	filename := fmt.Sprintf("data/day%vinput_sample.txt", day)
	input := readFileAsLines(filename)

	firstLine := true

	gridArray := GridComputers{}
lineLoop:
	for _, line := range input {
		if firstLine {
			firstLine = false
			continue lineLoop
		}
		//fmt.Printf("Line: %v\n", line)
		gc := NewGridComputer(line)
		gridArray = append(gridArray, gc)
		if maxX < gc.Loc.x {
			maxX = gc.Loc.x
		}
		if maxY < gc.Loc.y {
			maxY = gc.Loc.y
		}
		fmt.Println(gc.String())
	}
	//Max vals are based on 0-based values. Incr to reflect actual max
	maxX++
	maxY++
	ZeroPoint = &Point{x: 0, y: 0}

	// Part 1:
	validNodeCount := 0
	for _, gc := range gridArray {
		for _, other := range gridArray {
			if !gc.Equals(other) {
				if gc.Used > 0 {
					if gc.Used <= other.Avail {
						validNodeCount++
					}
				}
			}
		}
	}
	fmt.Printf("ValidNodeCount: %v\n", validNodeCount)

	// Part 2:
	// Construct the grid
	var cluster *GridComputerCluster
	var grid [][]*GridComputer
	fmt.Printf("MaxX: %v, MaxY: %v\n", maxX, maxY)
	grid = NewEmptyGrid()
	for _, gc := range gridArray {
		fmt.Printf("Putting %v at %v, %v\n", gc, gc.Loc.x, gc.Loc.y)
		grid[gc.Loc.y][gc.Loc.x] = gc
	}

	//fmt.Printf("%v\n", grid)
	goalPos := NewPoint(maxX-1, 0)
	fmt.Printf("Goal starting at %v\n", goalPos)

	cluster = &GridComputerCluster{
		Grid:      grid,
		Goal:      goalPos,
		MoveCount: 0,
	}

	foundBest := false
	best := 0
	seenClusterHash := make(map[string]bool)

	possibleClusters := GridComputerClusters{cluster}
	checks := 0
	fmt.Println("Starting BFS...\n")
	for len(possibleClusters) > 0 {
		nextCluster := possibleClusters[0]
		nextCluster.PrintCluster()
		checks++
		possibleClusters = possibleClusters[1:]
		newClusters := nextCluster.Permutations()
		//if checks > 4 {
		//	os.Exit(1)
		//}

		fmt.Printf("(%v)(%v)(%v)(%v) Goal now at [%v,%v] moves: %v, score: %v                  \n",
			checks, len(seenClusterHash), best, len(possibleClusters), nextCluster.Goal.x, nextCluster.Goal.y,
			nextCluster.MoveCount, nextCluster.Score())
		for _, cluster := range newClusters {
			//fmt.Printf("(%v)New Cluster:\n", len(possibleClusters))
			//cluster.PrintCluster()

			if cluster.Goal.x == 0 && cluster.Goal.y == 0 {
				if !foundBest || best > cluster.MoveCount {
					fmt.Printf("\nNewBest! %v\n", cluster.MoveCount)
					//os.Exit(1)
					best = cluster.MoveCount
					foundBest = true
					//os.Exit(1)
				}
			} else {
				if _, ok := seenClusterHash[cluster.Hash()]; !ok {
					//fmt.Printf("Appending new cluster %v=%v\n", cluster.Score(), cluster.Hash())
					if !foundBest || cluster.MoveCount < best {
						possibleClusters = append(possibleClusters, cluster)
						seenClusterHash[cluster.Hash()] = true
					} // else this is not as good as a found best, drop it.

				} else {
					//fmt.Printf("Already seen %v\n", cluster.Hash())
				}
			}
		}
		sort.Sort(possibleClusters)
	}
	fmt.Printf("\n(%v)(%v) Done with BFS\n", checks, best)
	fmt.Printf("Best: %v\n", best)
}

func NewEmptyGrid() [][]*GridComputer {
	grid := make([][]*GridComputer, maxY)
	for y := 0; y < maxY; y++ {
		grid[y] = make([]*GridComputer, maxX)
	}

	return grid
}

func (gcc *GridComputerCluster) CanBeMovedPairs() []*GridPair {
	if len(gcc.canBeMoved) > 0 {
		return gcc.canBeMoved
	}
	canBeMoved := []*GridPair{}
	for y := 0; y < len(gcc.Grid); y++ {
		for x := 0; x < len(gcc.Grid[0]); x++ {
			gc := gcc.Grid[y][x]
			gcPoint := gc.Loc
			if gc.Used > 0 { // Can't move empty data!
				moves := MovesForPoint(x, y)
				for _, move := range moves {
					toGc := gcc.Grid[move.y][move.x]
					if toGc.HasSpaceFor(gc, gcc.Goal) {
						//fmt.Printf("%v has space for %v\n", toGc, gc)
						canBeMoved = append(canBeMoved, &GridPair{
							A: gcPoint,
							B: move,
						})
					}
				}
			}
		}
	}
	//fmt.Printf("CanBeMoved: %v\n", canBeMoved)
	gcc.canBeMoved = canBeMoved
	return canBeMoved
}

func (gcc *GridComputerCluster) PrintCluster() {
	openCount := 0
	fmt.Printf("v---")
	for i := 0; i < maxX-2; i++ {
		fmt.Printf("----")
	}
	fmt.Println("---v")
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			gc := gcc.Grid[y][x]
			if gc.Loc.Equals(gcc.Goal) {
				fmt.Printf(" G  ")
			} else if gc.Loc.x == 0 && gc.Loc.y == 0 {
				aChar, uChar := gc.AvailUsedCharacters()
				fmt.Printf("(%v%v)", string(aChar), string(uChar))
			} else if gc.Used == 0 {
				fmt.Printf(" -- ")
				openCount++
			} else {
				aChar, uChar := gc.AvailUsedCharacters()
				fmt.Printf(" %s%s ", string(aChar), string(uChar))
			}
		}
		fmt.Println()
	}
	fmt.Printf("^---")
	for i := 0; i < maxX-2; i++ {
		fmt.Printf("----")
	}
	fmt.Println("---^")
	fmt.Printf("Open Count: %v, Move Count: %v, Score: %v\n", openCount, gcc.MoveCount, gcc.Score())
}

type GridPair struct {
	A *Point // from
	B *Point // to
}

func (gp *GridPair) String() string {
	return fmt.Sprintf("A: %v, B: %v", gp.A, gp.B)
}

func MovesForPoint(x int, y int) []*Point {
	//move left?
	potentialMoves := []*Point{}
	if x != 0 {
		potentialMoves = append(potentialMoves, NewPoint(x-1, y))
	}
	//move right?
	if x != maxX-1 {
		potentialMoves = append(potentialMoves, NewPoint(x+1, y))
	}
	//move up?
	if y != 0 {
		potentialMoves = append(potentialMoves, NewPoint(x, y-1))
	}
	//move down?
	if y != maxY-1 {
		potentialMoves = append(potentialMoves, NewPoint(x, y+1))
	}
	//fmt.Printf("Potential moves: %v\n", potentialMoves)
	return potentialMoves
}

func (gcc *GridComputerCluster) Permutations() []*GridComputerCluster {

	gccs := []*GridComputerCluster{}
	validDataMovePairs := gcc.CanBeMovedPairs()
	//fmt.Printf("Valid Pairs: %v\n", validDataMovePairs)

	for _, pair := range validDataMovePairs {
		gccs = append(gccs, gcc.MoveData(pair))

	}
	return gccs
}

func (gcc *GridComputerCluster) MoveData(pair *GridPair) *GridComputerCluster {

	//fmt.Printf("Moving data from %v to %v\n", pair.A, pair.B)
	newGrid := NewEmptyGrid()
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			curGc := gcc.Grid[y][x]
			newGrid[y][x] = curGc.Copy()
		}
	}

	fromGc := newGrid[pair.A.y][pair.A.x]
	toGc := newGrid[pair.B.y][pair.B.x]

	// Move actual data:
	//fmt.Printf("Moving from %v to %v\n", fromGc, toGc)
	toGc.Avail -= fromGc.Used
	toGc.Used += fromGc.Used
	fromGc.Avail += fromGc.Used
	fromGc.Used = 0
	//fmt.Printf("From now %v, To now %v\n", fromGc, toGc)

	if toGc.Avail < 0 {
		fmt.Printf("WTF? from: %v, to: %v, Grid:%v\n", fromGc, toGc, newGrid)
		os.Exit(1)
	}

	// Adjust goal, if we moved data from the Goal
	var goal *Point
	if gcc.Goal.x == fromGc.Loc.x && gcc.Goal.y == fromGc.Loc.y {
		// Also indicate that we reached the goal, so we can change score weights
		gcc.goalReached = true

		goal = &Point{
			x: toGc.Loc.x,
			y: toGc.Loc.y,
		}
		//fmt.Printf("Adjusting Goal location to %v\n", goal)
	} else {
		goal = &Point{
			x: gcc.Goal.x,
			y: gcc.Goal.y,
		}
	}

	return &GridComputerCluster{
		Grid:        newGrid,
		Goal:        goal,
		MoveCount:   gcc.MoveCount + 1,
		goalReached: gcc.goalReached,
	}
}

type GridComputerClusters []*GridComputerCluster

type GridComputerCluster struct {
	Grid        [][]*GridComputer
	Goal        *Point
	canBeMoved  []*GridPair
	MoveCount   int
	hashValue   string
	goalReached bool
}

func (gcc *GridComputerCluster) Score() int {
	// Lowest score best. Sort will order lowest to highest.

	//Find the closest 'FromGC' point, and use that distance in the score.
	dist := 9999
	moveWeight := 3
	for _, pair := range gcc.CanBeMovedPairs() {
		zeroPointWeight := 1
		goalWeight := 1

		if pair.A.y > 11 {
			zeroPointWeight = 10
		}
		if gcc.goalReached {
			zeroPointWeight = 5 // trend toward zero
			goalWeight = 8      // stick close to goal
			moveWeight = 1
		} else {
			goalWeight = 7
		}
		nextScore := zeroPointWeight * pair.A.DistanceBetween(ZeroPoint)
		nextScore += goalWeight * pair.A.DistanceBetween(gcc.Goal) // Make it hone in on Goal first
		if nextScore < dist {
			dist = nextScore
		}
	}
	//fmt.Printf("Calculated score for %v as %d\n", gcc.Hash(), dist+gcc.MoveCount)
	return dist + moveWeight*gcc.MoveCount

	// Old: Add scores of all points that can be moved
	//score := 0
	//for _, pair := range gcc.CanBeMovedPairs() {
	//	score += pair.A.DistanceBetween(ZeroPoint)
	//	score += pair.A.DistanceBetween(gcc.Goal)
	//}

	//return score + gcc.Goal.moveCount
}

func (gcc *GridComputerCluster) String() string {
	return gcc.Hash()
}

func (gcc *GridComputerCluster) Hash() string {
	if gcc.hashValue == "" {
		toHash := ""

		for y := 0; y < len(gcc.Grid); y++ {
			for x := 0; x < len(gcc.Grid[0]); x++ {
				toHash += gcc.Grid[y][x].String()
			}
		}
		toHash += gcc.Goal.String()
		gcc.hashValue = fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
	}
	return gcc.hashValue
}

func (p GridComputerClusters) Len() int { return len(p) }
func (p GridComputerClusters) Less(i, j int) bool {
	return p[i].Score() < p[j].Score()
}
func (p GridComputerClusters) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type Point struct {
	x int
	y int
}

func (p *Point) String() string {
	return fmt.Sprintf("[%v,%v]", p.x, p.y)
}

func (p *Point) DistanceBetween(other *Point) int {
	xDist := p.x - other.x
	yDist := p.y - other.y
	return int(math.Abs(float64(xDist)) + math.Abs(float64(yDist)))
}

func (p *Point) Equals(other *Point) bool {
	return p.x == other.x && p.y == other.y
}

func NewPoint(x int, y int) *Point {
	return &Point{x: x, y: y}
}

type GridComputers []*GridComputer

type GridComputer struct {
	Loc   *Point
	Used  int
	Avail int
}

func (gc *GridComputer) AvailUsedCharacters() (rune, rune) {
	a := gc.Avail
	u := gc.Used
	var aChar rune
	var uChar rune
	switch {
	case a < 10:
		aChar = '.'
	case a >= 10 && a < 20:
		aChar = 'a'
	case a >= 20 && a < 30:
		aChar = 'A'
	case a >= 30:
		aChar = '@'
	}
	switch {
	case u < 25:
		uChar = '.'
	case u >= 25 && u < 75:
		uChar = 'o'
	case u >= 75 && u < 100:
		uChar = 'U'
	case u >= 100:
		uChar = '#'
	}
	return aChar, uChar
}

func (gc *GridComputer) Equals(other *GridComputer) bool {
	return gc.Loc.Equals(other.Loc)
}

func (gc *GridComputer) HasSpaceFor(other *GridComputer, goal *Point) bool {
	if other.Loc.x == goal.x && other.Loc.y == goal.y {
		// There must be room for all of it, if we're moving the Goal data
		//fmt.Printf("Trying to move Goal Data. Can we? %v\n", gc.Used == 0)
		return gc.Used == 0
	} else {
		return gc.Avail >= other.Used
	}
}

func (gc *GridComputer) String() string {
	return fmt.Sprintf("[x%v-y%v]\tU: %v\tA: %v", gc.Loc.x, gc.Loc.y, gc.Used, gc.Avail)
}

func (gc *GridComputer) Copy() *GridComputer {
	return &GridComputer{
		Loc: &Point{
			x: gc.Loc.x,
			y: gc.Loc.y,
		},
		Used:  gc.Used,
		Avail: gc.Avail,
	}
}

func NewGridComputer(line string) *GridComputer {

	// /dev/grid/node-x0-y0     88T   67T    21T   76%

	tokens := strings.Split(line, " ")
	nonEmptyTokens := []string{}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] != "" {
			nonEmptyTokens = append(nonEmptyTokens, tokens[i])
		}
	}

	//fmt.Printf("Tokens: %v\n", nonEmptyTokens)
	machine := nonEmptyTokens[0]
	prefix := "/dev/grid/node-"
	node := machine[len(prefix):]
	nodeTokens := strings.Split(node, "-")
	//fmt.Printf("Node tokens: %v\n", nodeTokens)

	Xpos, _ := strconv.Atoi(nodeTokens[0][1:])
	Ypos, _ := strconv.Atoi(nodeTokens[1][1:])
	Used, _ := strconv.Atoi(nonEmptyTokens[2][0 : len(nonEmptyTokens[2])-1])
	Avail, _ := strconv.Atoi(nonEmptyTokens[3][0 : len(nonEmptyTokens[3])-1])
	return &GridComputer{
		Loc: &Point{
			x: Xpos,
			y: Ypos,
		},
		Used:  Used,
		Avail: Avail,
	}
}
