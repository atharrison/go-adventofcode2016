package adventofcode

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var maxX int
var maxY int

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
		if maxX < gc.Xpos {
			maxX = gc.Xpos
		}
		if maxY < gc.Ypos {
			maxY = gc.Ypos
		}
		fmt.Println(gc.String())
	}
	//Max vals are based on 0-based values. Incr to reflect actual
	maxX++
	maxY++

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
		fmt.Printf("Putting %v at %v, %v\n", gc, gc.Xpos, gc.Ypos)
		grid[gc.Ypos][gc.Xpos] = gc
	}

	//fmt.Printf("%v\n", grid)
	goalPos := NewPoint(0, maxY-1)

	cluster = &GridComputerCluster{
		Grid: grid,
		Goal: goalPos,
	}

	foundBest := false
	best := 0
	seenClusterHash := make(map[string]bool)

	possibleClusters := GridComputerClusters{cluster}
	checks := 0
	fmt.Println("Starting BFS...\n")
	for len(possibleClusters) > 0 {
		nextCluster := possibleClusters[0]
		checks++
		possibleClusters := possibleClusters[1:]
		newClusters := nextCluster.Permutations()

		fmt.Printf("(%v)(%v)(%v) Goal now at [%v,%v] moves: %v                  \r", checks, best, len(possibleClusters), nextCluster.Goal.x, nextCluster.Goal.y, nextCluster.Goal.moveCount)
		for _, cluster := range newClusters {
			if cluster.Goal.x == 0 && cluster.Goal.y == 0 {
				if !foundBest || best > cluster.Goal.moveCount {
					fmt.Printf("NewBest! %v\n", cluster.Goal.moveCount)
					best = cluster.Goal.moveCount
				}
			} else {
				if _, ok := seenClusterHash[cluster.Hash()]; !ok {
					possibleClusters = append(possibleClusters, cluster)
					seenClusterHash[cluster.Hash()] = true
				} else {
					fmt.Printf("Already seen %v\n", cluster.Hash())
				}
				//perms := nextCluster.Permutations()
				//fmt.Printf("Next cluster has %v possible permutations\n", len(perms))
				//for _, c := range perms {
				//
				//}
			}
		}
		sort.Sort(possibleClusters)
	}
	fmt.Printf("(%v)(%v) Done with BFS\n", checks, best)
}

func NewEmptyGrid() [][]*GridComputer {
	grid := make([][]*GridComputer, maxY)
	for y := 0; y < maxY; y++ {
		grid[y] = make([]*GridComputer, maxX)
	}

	return grid
}

func (gcc *GridComputerCluster) CanBeMovedPairs() []*GridPair {
	canBeMoved := []*GridPair{}
	for y := 0; y < len(gcc.Grid); y++ {
		for x := 0; x < len(gcc.Grid[0]); x++ {
			gc := gcc.Grid[y][x]
			gcPoint := &Point{x: gc.Xpos, y: gc.Ypos}
			if gc.Used > 0 { // Can't move empty data!
				moves := MovesForPoint(x, y)
				for _, move := range moves {
					toGc := gcc.Grid[move.y][move.x]
					if toGc.HasSpaceFor(gc) {
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
	return canBeMoved
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

	fmt.Printf("Moving data from %v to %v\n", pair.A, pair.B)
	newGrid := NewEmptyGrid()
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			curGc := gcc.Grid[y][x]
			newGrid[y][x] = curGc.Copy()
		}
	}

	fromGc := gcc.Grid[pair.A.y][pair.A.x]
	toGc := gcc.Grid[pair.B.y][pair.B.x]

	// Move actual data:
	toGc.Avail -= fromGc.Used
	fromGc.Used = 0

	// Adjust goal, if we moved data from the Goal
	var goal *Point
	if gcc.Goal.x == fromGc.Xpos && gcc.Goal.y == fromGc.Ypos {
		goal = &Point{
			x:         toGc.Xpos,
			y:         toGc.Ypos,
			moveCount: gcc.Goal.moveCount + 1,
		}

	} else {
		goal = gcc.Goal
	}

	return &GridComputerCluster{
		Grid: newGrid,
		Goal: goal,
	}
}

type GridComputerClusters []*GridComputerCluster

type GridComputerCluster struct {
	Grid      [][]*GridComputer
	Goal      *Point
	hashValue string
}

func (gcc *GridComputerCluster) Score() int {
	// Lowest score best?
	return gcc.Goal.y + gcc.Goal.moveCount
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
	x         int
	y         int
	moveCount int
}

func (p *Point) String() string {
	return fmt.Sprintf("[%v,%v](%v)", p.x, p.y, p.moveCount)
}

func NewPoint(x int, y int) *Point {
	return &Point{x: x, y: y, moveCount: 0}
}

type GridComputers []*GridComputer

type GridComputer struct {
	Xpos  int
	Ypos  int
	Used  int
	Avail int
}

func (gc *GridComputer) Equals(other *GridComputer) bool {
	return gc.Xpos == other.Xpos && gc.Ypos == other.Ypos
}

func (gc *GridComputer) HasSpaceFor(other *GridComputer) bool {
	return gc.Avail >= other.Used
}

func (gc *GridComputer) String() string {
	return fmt.Sprintf("[x%v-y%v]\tU: %v\tA: %v", gc.Xpos, gc.Ypos, gc.Used, gc.Avail)
}

func (gc *GridComputer) Copy() *GridComputer {
	return &GridComputer{
		Xpos:  gc.Xpos,
		Ypos:  gc.Ypos,
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
		Xpos:  Xpos,
		Ypos:  Ypos,
		Used:  Used,
		Avail: Avail,
	}
}
