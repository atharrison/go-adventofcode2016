package adventofcode

import (
	"crypto/md5"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var ductMap [][]*DuctPoint
var DuctGoals map[int]*DuctPoint
var robotsSeenHash map[string]bool

var ductWidth int
var ductHeight int

const KillEarly = false

func (r *DuctRobot) PrintRobot(best int, possibleRobotsLen int) {
	//fmt.Printf("                                                                 \n")
	fmt.Printf("(%v)(%v)(%v) [%v, %v]\tMoves: %v\tMet: %v\tScore: %v            \n", best, possibleRobotsLen, len(robotsSeenHash), r.X, r.Y, r.MoveCount, len(r.GoalsMet), r.Score())
}

func Day24() {
	day := "24"
	filename := fmt.Sprintf("data/day%vinput_sample.txt", day)
	//filename := fmt.Sprintf("data/day%vinput.txt", day)
	input := readFileAsLines(filename)

	robotsSeenHash = make(map[string]bool)

	DuctGoals = make(map[int]*DuctPoint)

	var StartPoint *DuctPoint

	ductWidth = len(input[0])
	ductHeight = len(input)
	ductMap = make([][]*DuctPoint, ductHeight)

	fmt.Printf("Height: %v, Width: %v\n", ductHeight, ductWidth)
	for y, line := range input {
		fmt.Printf("Line: %v\n", line)
		ductMap[y] = make([]*DuctPoint, ductWidth)
		for x, c := range line {

			if c == '0' {
				StartPoint = &DuctPoint{X: x, Y: y, T: c}
				c = '.'
			}

			ductMap[y][x] = &DuctPoint{X: x, Y: y, T: c}
			if isGoalPoint(c) {
				goalInt, _ := strconv.Atoi(string(c))
				DuctGoals[goalInt] = &DuctPoint{X: x, Y: y, T: c}
			}
		}
	}

	robot := &DuctRobot{
		X:         StartPoint.X,
		Y:         StartPoint.Y,
		MoveCount: 0,
		GoalsMet:  make(map[int]bool),
		Path:      NewDuctPath(),
	}
	robotsSeenHash[robot.Hash()] = true

	fmt.Printf("Goals: %v\n", DuctGoals)
	fmt.Printf("Start: %v\n", StartPoint)

	MoveRobotThroughDuct(robot)

	//fmt.Printf("%v\n", ductMap)
}

func isGoalPoint(c rune) bool {
	return c >= 49 && c <= 57
}

type DuctRobot struct {
	X         int
	Y         int
	MoveCount int
	GoalsMet  map[int]bool
	Path      [][]bool
}

func NewDuctPath() [][]bool {
	path := make([][]bool, ductHeight)
	for y := 0; y < ductHeight; y++ {
		path[y] = make([]bool, ductWidth)
	}
	return path
}

func CopyDuctPath(existing [][]bool) [][]bool {
	path := make([][]bool, ductHeight)
	for y := 0; y < ductHeight; y++ {
		path[y] = make([]bool, ductWidth)
		for x := 0; x < ductWidth; x++ {
			path[y][x] = existing[y][x]
		}
	}
	return path

}

func NewDuctRobot(x int, y int, goalsMet map[int]bool, moveCount int, existing [][]bool) *DuctRobot {
	path := CopyDuctPath(existing)
	path[y][x] = true
	return &DuctRobot{
		X:         x,
		Y:         y,
		GoalsMet:  goalsMet,
		MoveCount: moveCount,
		Path:      path,
	}
}

func (r *DuctRobot) Hash() string {
	toHash := ""
	for g := 1; g < len(DuctGoals)+1; g++ {
		if _, ok := r.GoalsMet[g]; ok {
			toHash += fmt.Sprintf("%v", g)
		}
	}

	for y := 0; y < ductHeight; y++ {
		for x := 0; x < ductWidth; x++ {
			var val string
			if r.Path[y][x] {
				val = "T"
			} else {
				val = "F"
			}
			toHash += val
		}
	}
	fmt.Printf("Robot Fingerprint: %v\n", toHash)
	return fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
}

type DuctMap [][]*DuctPoint

type DuctRobots []*DuctRobot

func (p DuctRobots) Len() int { return len(p) }
func (p DuctRobots) Less(i, j int) bool {
	return p[i].Score() < p[j].Score()
}
func (p DuctRobots) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type DuctPoint struct {
	T rune
	X int
	Y int
}

func (p *DuctPoint) String() string {
	//return fmt.Sprintf("%s", string(p.T))
	return fmt.Sprintf("%v[%v, %v]", string(p.T), p.X, p.Y)
}

func MoveRobotThroughDuct(robot *DuctRobot) {

	possibleRobots := DuctRobots{robot}
	best := 9999999
	var bestPath [][]bool
	bestReached := false

	totalMoves := 0

MainLoop:
	for len(possibleRobots) > 0 {
		nextRobot := possibleRobots[0]
		possibleRobots = possibleRobots[1:]
		totalMoves++

		if KillEarly && totalMoves > 5 {
			os.Exit(1)
		}

		nextRobot.PrintRobot(best, len(possibleRobots))

		if nextRobot.AllGoalsMet() {
			fmt.Printf("All Goals Met: %v\n", nextRobot.MoveCount)
			if !bestReached || nextRobot.MoveCount < best {
				best = nextRobot.MoveCount
				bestReached = true
				bestPath = nextRobot.Path
			}
			continue MainLoop // Done with this robot, success.
		} else if bestReached && nextRobot.MoveCount > best {
			continue MainLoop // Other robots have done better, retire.
		}

		newRobots := nextRobot.PossibleMoves()
		for _, newRobot := range newRobots {
			robotHash := newRobot.Hash()
			if _, ok := robotsSeenHash[robotHash]; !ok {
				possibleRobots = append(possibleRobots, newRobot)
				robotsSeenHash[robotHash] = true
			} else {
				//fmt.Println("Already seen %v\n", newRobot.Hash())
				//Exclude this movement, we've done it before.
			}
		}

		sort.Sort(possibleRobots)

	}

	fmt.Printf("Finished. Best: %v\n", best)
	for y := 0; y < ductHeight; y++ {
		for x := 0; x < ductWidth; x++ {
			if bestPath[y][x] {
				fmt.Printf("=")
			} else {
				fmt.Printf("x")
			}
		}
		fmt.Println()
	}

}

func (r *DuctRobot) PossibleMoves() []*DuctRobot {
	potentialMoves := []*DuctPoint{}
	newRobots := []*DuctRobot{}

	//fmt.Printf("Checking Robot %v for Possible Moves.\n", r)
	//move left?
	if r.X != 0 {
		potentialMoves = append(potentialMoves, ductMap[r.Y][r.X-1])
	}
	//move right?
	if r.X != ductWidth-1 {
		potentialMoves = append(potentialMoves, ductMap[r.Y][r.X+1])
	}
	////move up?
	if r.Y != 0 {
		potentialMoves = append(potentialMoves, ductMap[r.Y-1][r.X])
	}
	////move down?
	if r.Y != ductHeight-1 {
		potentialMoves = append(potentialMoves, ductMap[r.Y+1][r.X])
	}

	for _, move := range potentialMoves {
		if move.T != '#' {
			newGoals := r.CopyGoals()
			if isGoalPoint(move.T) {
				goalInt, _ := strconv.Atoi(string(move.T))
				newGoals[goalInt] = true
				fmt.Printf("\nFound Goal %v\n", goalInt)
			}
			newRobots = append(newRobots, NewDuctRobot(move.X, move.Y, newGoals, r.MoveCount+1, CopyDuctPath(r.Path)))
		}
	}
	//fmt.Printf("Potential moves: %v\n", potentialMoves)
	//return potentialMoves
	return newRobots
}

func (r *DuctRobot) CopyGoals() map[int]bool {
	newGoals := make(map[int]bool)
	for k, v := range r.GoalsMet {
		newGoals[k] = v
	}
	return newGoals
}

func (r *DuctRobot) AllGoalsMet() bool {
	return len(r.GoalsMet) == len(DuctGoals)
}

func (r *DuctRobot) Score() int {

	// Negative points for not meeting goals
	score := 500 * (len(DuctGoals) - len(r.GoalsMet))

	if len(r.GoalsMet) < len(DuctGoals) {
		closest := 999999
		for goalInt, v := range DuctGoals {
			if !r.GoalsMet[goalInt] { // Don't care about dist to goals already met
				//score += v.DistanceToRobot(r)
				//score += goalInt * v.DistanceToRobot(r)
				nextDist := v.DistanceToRobot(r)
				if closest > nextDist {
					closest = nextDist
				}
			}
		}
		score += closest // Go toward closet point
	}

	return score

}

func (p *DuctPoint) DistanceToRobot(r *DuctRobot) int {
	xDist := p.X - r.X
	yDist := p.Y - r.Y
	return int(math.Abs(float64(xDist)) + math.Abs(float64(yDist)))
}

/////////////////////////////////////////////////////////
/*
New tactic: Find distances from each point a,b. Then solve the Traveling Salesman problem for those points/distances.

Copy 'WalkFloor()' from Day13, but plug in the duct map from Day 24, and provide start points.
*/

func Day24Alt() {
	day := "24"
	//filename := fmt.Sprintf("data/day%vinput_sample.txt", day)
	filename := fmt.Sprintf("data/day%vinput.txt", day)
	input := readFileAsLines(filename)

	robotsSeenHash = make(map[string]bool)

	DuctGoals = make(map[int]*DuctPoint)

	ductWidth = len(input[0])
	ductHeight = len(input)
	ductMapStrings := make([][]string, ductHeight)

	fmt.Printf("Height: %v, Width: %v\n", ductHeight, ductWidth)
	for y, line := range input {
		fmt.Printf("Line: %v\n", line)
		ductMapStrings[y] = make([]string, ductWidth)
		for x, c := range line {

			if c == '0' || isGoalPoint(c) {
				goalInt, _ := strconv.Atoi(string(c))
				DuctGoals[goalInt] = &DuctPoint{X: x, Y: y, T: c}
				c = '.'
				// 'Start Point' will be DuctGoals[0]
			}

			ductMapStrings[y][x] = string(c)
		}
	}

	fmt.Printf("Goals: %v\n", DuctGoals)

	GoalDistances := make([][]int, len(DuctGoals))
	for i := 0; i < len(DuctGoals); i++ {
		GoalDistances[i] = make([]int, len(DuctGoals))
	}

	fmt.Printf("%v\n", ductMap)
	for a, p1 := range DuctGoals {
		for b := a + 1; b < len(DuctGoals); b++ {
			p2 := DuctGoals[b]

			dist := WalkDuctMap(ductMapStrings, p1.X, p1.Y, p2.X, p2.Y)
			fmt.Printf("Processed Goal %v -> Goal %v Dist: %v\n", p1, p2, dist)
			GoalDistances[a][b] = dist
			GoalDistances[b][a] = dist
		}
	}

	fmt.Printf("GoalDistances:\n%v\n", GoalDistances)
	fmt.Println("   0\t1\t2\t3\t4\t5\t6\t7")
	for a := 0; a < len(DuctGoals); a++ {
		fmt.Printf("%v: ", a)
		for b := 0; b < len(DuctGoals); b++ {
			fmt.Printf("%v\t", GoalDistances[a][b])
		}
		fmt.Println()
	}

	//Now, we have distances from each point to each other point.
	//Solve the traveling salesman problem, starting at 0, touching every other point, but not returning to 0.
	//Smallest distance will be our answer.

	shortestDistance, shortestPath := SolveModifiedTraveler(GoalDistances)
	fmt.Printf("Shortest distance: %v, Path: %v\n", shortestDistance, shortestPath)
}

func SolveModifiedTraveler(distances [][]int) (int, *TravelerPath) {

	numGoals := len(distances[0])
	bestDist := 9999
	var bestPath *TravelerPath

	travelerPath := &TravelerPath{
		Reached:   []int{0},
		TotalDist: 0,
	}

	traveledPaths := TraveledPaths{travelerPath}

	checks := 0
	for len(traveledPaths) > 0 {
		nextPath := traveledPaths[0]
		traveledPaths = traveledPaths[1:]

		//if checks > 5 {
		//	os.Exit(1)
		//}
		if len(nextPath.Reached) == numGoals {
			if nextPath.TotalDist < bestDist {
				bestDist = nextPath.TotalDist
				bestPath = nextPath
			}
			fmt.Printf("Found best: %v\n", bestDist)
		} else if nextPath.TotalDist > bestDist {
			//Too far, bail by not adding combos from here.
		} else {
			//Pick next best distance:
			for b := 1; b < numGoals; b++ {
				if !nextPath.HasReachedGoal(b) {

					newReached := []int{}
					for _, j := range nextPath.Reached {
						newReached = append(newReached, j)
					}
					newReached = append(newReached, b)

					newPath := &TravelerPath{
						CurrentLoc: b,
						Reached:    newReached,
						TotalDist:  nextPath.TotalDist + distances[nextPath.CurrentLoc][b],
					}

					traveledPaths = append(traveledPaths, newPath)
					fmt.Printf("(%v) New Path: %v\n", len(traveledPaths), newPath)
				}
			}
		}
		sort.Sort(traveledPaths)
		checks++
	}

	return bestDist, bestPath
}

type TraveledPaths []*TravelerPath

func (slice TraveledPaths) Len() int {
	return len(slice)
}
func (slice TraveledPaths) Less(i, j int) bool {
	return slice[i].TotalDist < slice[j].TotalDist
}
func (slice TraveledPaths) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

type TravelerPath struct {
	CurrentLoc int
	Reached    []int
	TotalDist  int
}

func (p *TravelerPath) HasReachedGoal(goal int) bool {
	for _, g := range p.Reached {
		if g == goal {
			return true
		}
	}
	return false
}

func WalkDuctMap(floorplan [][]string, startX, startY, goalX, goalY int) int {

	lowestDist := 9999
	stepsToProcess := Locations{
		&Location{
			X:    startX,
			Y:    startY,
			Dist: 0,
		},
	}

	var ductTraveledHash = make(map[string]*Location)

	for len(stepsToProcess) > 0 {

		next := stepsToProcess[0]
		stepsToProcess = stepsToProcess[1:]
		//fmt.Printf("Best: %v\t ToProcess: %v\tNow at [%v, %v] D:%v    \n", lowestDist, len(stepsToProcess), next.X, next.Y, next.Dist)
		if next.Dist > lowestDist {
			continue
		}

		if next.X == goalX && next.Y == goalY && lowestDist > next.Dist {
			lowestDist = next.Dist
			//fmt.Printf("Best Dist now %v\n", lowestDist)
		}

		for xDelta := -1; xDelta < 2; xDelta++ {
			if next.X+xDelta < 0 || next.X+xDelta >= ductWidth {
				continue
			}
			for yDelta := -1; yDelta < 2; yDelta++ {
				if next.Y+yDelta < 0 || next.Y+yDelta >= ductHeight {
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
					if loc, ok := ductTraveledHash[h]; ok {
						if loc.Dist > newLoc.Dist {
							ductTraveledHash[h] = newLoc
							stepsToProcess = append(stepsToProcess, newLoc)
						}
					} else {
						stepsToProcess = append(stepsToProcess, newLoc)
						ductTraveledHash[h] = newLoc
					}

				}
			}
		}
		sort.Sort(stepsToProcess)
	}

	return lowestDist
}
