package adventofcode

import (
	"crypto/md5"
	"fmt"
	"sort"
)

var bestTotalMoves int
var maxMicrochips int
var maxGenerators int
var maxFloor int
var totalCombinationsCalculated int

var floorCombinationHash map[string]bool

func Day11() {
	//day := "11"
	//filename := fmt.Sprintf("data/day%vinput.txt", day)
	//input := readFileAsLines(filename)
	//
	//for _, line := range input {
	//	fmt.Printf("Line: %v\n", line)
	//}
	//fmt.Printf("%v\n", input)

	maxFloor = 4
	bestTotalMoves = 9999999
	var floorMap map[int]*RTGFloor

	/*
		F4 . .  .  .  .  .  .  .  .  .  .
		F3 . .  .  PG PM RG RM .  .  .  .
		F2 . .  LM .  .  .  .  .  SM .  .
		F1 E LG .  .  .  .  .  SG .  TG TM
	*/
	maxMicrochips = 5
	maxGenerators = 5
	floorMap = map[int]*RTGFloor{
		1: &RTGFloor{
			Number:     1,
			Generators: []string{"LG", "SG", "TG"},
			Microchips: []string{"TM"},
		},
		2: &RTGFloor{
			Number:     2,
			Generators: []string{},
			Microchips: []string{"LM", "SM"},
		},
		3: &RTGFloor{
			Number:     3,
			Generators: []string{"PG", "RG"},
			Microchips: []string{"PM", "RM"},
		},
		4: &RTGFloor{
			Number:     4,
			Generators: []string{},
			Microchips: []string{},
		},
	}

	var floorMapSample = map[int]*RTGFloor{
		1: &RTGFloor{
			Number:     1,
			Generators: []string{},
			Microchips: []string{"HM", "LM"},
		},
		2: &RTGFloor{
			Number:     2,
			Generators: []string{"HG"},
			Microchips: []string{},
		},
		3: &RTGFloor{
			Number:     3,
			Generators: []string{"LG"},
			Microchips: []string{},
		},
		4: &RTGFloor{
			Number:     4,
			Generators: []string{},
			Microchips: []string{},
		},
	}

	// Sample setup:
	if false {
		maxMicrochips = 2
		maxGenerators = 2
		floorMap = floorMapSample
	}
	// End S

	var elevator = &RTGElevator{
		Floor:      1,
		Generators: []string{},
		Microchips: []string{},
	}
	floorCombinationHash = make(map[string]bool)

	floorCombinationHash[FloorMapAndElevatorHash(floorMap, elevator)] = true

	//numMoves := MoveRTGItems(floorMap, elevator, 0)
	StartRTGMoves(floorMap, elevator)
	//fmt.Printf("Moves: %v\n", numMoves)
	fmt.Printf("BestMoves: %v\n", bestTotalMoves)
}

func FloorMapAndElevatorHash(floorMap map[int]*RTGFloor, elevator *RTGElevator) string {
	toHash := ""
	for i := 1; i <= maxFloor; i++ {
		toHash += floorMap[i].String()
	}
	toHash += elevator.String()
	hash := fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
	return hash
}

type RTGElevator struct {
	Floor      int
	Generators []string
	Microchips []string
}

func (e *RTGElevator) String() string {
	return fmt.Sprintf("E%v with %v and %v", e.Floor, e.Generators, e.Microchips)
}

type RTGFloor struct {
	Number     int
	Generators []string
	Microchips []string
}

func (f *RTGFloor) String() string {

	return fmt.Sprintf("F%v has %v and %v", f.Number, f.Generators, f.Microchips)
}

func (f *RTGFloor) Score() int {
	return 2 * f.Number * (len(f.Generators) + len(f.Microchips))
}

func (f *RTGFloor) Items() []string {
	var items []string

	for _, g := range f.Generators {
		items = append(items, g)
	}
	for _, m := range f.Microchips {
		items = append(items, m)
	}
	return items
}

type MapAndMove struct {
	Elevator *RTGElevator
	FloorMap map[int]*RTGFloor
	Moves    int
}

func (mam *MapAndMove) Score() int {
	//return mam.Moves - FloorMapScore(mam.FloorMap)
	return FloorMapScore(mam.FloorMap)
}

func FloorMapScore(floorMap map[int]*RTGFloor) int {
	score := 0
	for _, f := range floorMap {
		score += f.Score()
	}
	return score
}

type MapAndMoveList []*MapAndMove

func (slice MapAndMoveList) Len() int {
	return len(slice)
}

func (slice MapAndMoveList) Less(i, j int) bool {
	return slice[i].Score() > slice[j].Score()
	//if slice[i].Moves == slice[j].Moves {
	//	return FloorMapScore(slice[i].FloorMap) < FloorMapScore(slice[j].FloorMap)
	//}
	//return slice[i].Moves < slice[j].Moves;
}

func (slice MapAndMoveList) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func StartRTGMoves(startFloorMap map[int]*RTGFloor, startElevator *RTGElevator) int {

	moves := 0
	possibleMoves := MapAndMoveList{
		&MapAndMove{
			Elevator: startElevator,
			FloorMap: startFloorMap,
			Moves:    0,
		},
	}
	for len(possibleMoves) > 0 {
		nextMove := possibleMoves[0]
		//fmt.Printf("%v (%v) Moves to look at\n", len(possibleMoves), nextMove.Moves)

		possibleMoves = possibleMoves[1:]
		elevator := nextMove.Elevator
		floorMap := nextMove.FloorMap

		options := GetOptionsForElevatorAndFloor(elevator, floorMap)
		validOptions := ValidOptions(options, floorMap, moves)
		for _, option := range validOptions {
			nextMapAndMove := GenerateNextFloorIteration(option, floorMap, nextMove.Moves)
			if nextMapAndMove != nil && !TopFloorHasEverything(nextMapAndMove.FloorMap[maxFloor]) {
				possibleMoves = append(possibleMoves, nextMapAndMove)
			}
		}
		sort.Sort(possibleMoves)
	}

	return bestTotalMoves
}

func ValidOptions(options []*RTGOptions, floorMap map[int]*RTGFloor, moves int) []*RTGOptions {
	validOptions := []*RTGOptions{}
	for _, option := range options {
		_, valid1 := NewFloorWithItems(option.Items, floorMap[option.GotoFloor])
		_, valid2 := NewFloorWithoutItems(option.Items, floorMap[option.Floor])
		if valid1 && valid2 {
			//if TopFloorHasEverything(newFloor) {
			//	if moves < bestTotalMoves {
			//		fmt.Printf("\nFound Possible best! %v\n", moves)
			//		bestTotalMoves = moves
			//		//return moves
			//	}
			//} else {
			validOptions = append(validOptions, option)
			//}
		}
	}
	return validOptions
}

func GenerateNextFloorIteration(option *RTGOptions, floorMap map[int]*RTGFloor, moves int) *MapAndMove {

	newFloor, _ := NewFloorWithItems(option.Items, floorMap[option.GotoFloor])
	oldFloor, _ := NewFloorWithoutItems(option.Items, floorMap[option.Floor])

	moves += 1
	if moves > 125 {
		return nil //I know it is not higher than 123, for Part1
	}
	if option.GotoFloor == maxFloor && TopFloorHasEverything(newFloor) {
		fmt.Printf("\nFound Possible best... %v\n", moves)
		if moves < bestTotalMoves {
			fmt.Printf("\nYes, New best! %v\n", moves)
			bestTotalMoves = moves
		} else {
			fmt.Printf("Complete but not best (%v), with %v moves.\n", bestTotalMoves, moves)
		}
		return nil
	}

	newFloorMap := MakeNewFloorMap(floorMap, newFloor, oldFloor)
	newElevator := &RTGElevator{
		Floor:      option.GotoFloor,
		Generators: option.Generators(),
		Microchips: option.Microchips(),
	}

	newFloorAndElevatorHash := FloorMapAndElevatorHash(newFloorMap, newElevator)
	if _, ok := floorCombinationHash[newFloorAndElevatorHash]; ok {
		//fmt.Printf("Already seen %v, skipping.\n", newFloorHash)
		return nil
	} else {
		floorCombinationHash[newFloorAndElevatorHash] = true
		totalCombinationsCalculated++
	}
	fmt.Printf("(%v)(%v)New E: {%v} FMS:[%v] %v\r", moves, totalCombinationsCalculated, newElevator, FloorMapScore(floorMap), newFloorMap)

	return &MapAndMove{
		Elevator: newElevator,
		FloorMap: newFloorMap,
		Moves:    moves,
	}
}

//func MoveRTGItems(floorMap map[int]*RTGFloor, elevator *RTGElevator, moves int) int {
//
//	moves = moves+1
//	if moves > bestTotalMoves {
//		//No need to continue
//		return bestTotalMoves
//	}
//
//	options := GetOptionsForElevatorAndFloor(elevator, floorMap)
//	//fmt.Printf("Options: %v\n", options)
//
//	validOptions := []*RTGOptions{}
//	for _, option := range options {
//		newFloor, valid1 := NewFloorWithItems(option.Items, floorMap[option.GotoFloor])
//		_, valid2 := NewFloorWithoutItems(option.Items, floorMap[option.Floor])
//		if valid1 && valid2 {
//			if TopFloorHasEverything(newFloor) {
//				if moves < bestTotalMoves {
//					fmt.Printf("\nFound Possible best! %v\n", moves)
//					bestTotalMoves = moves
//					return moves
//				}
//			} else {
//				validOptions = append(validOptions, option)
//			}
//		}
//	}
//
//	best := 9999999
//	for _, option := range validOptions {
//		newFloor, _ := NewFloorWithItems(option.Items, floorMap[option.GotoFloor])
//		oldFloor, _ := NewFloorWithoutItems(option.Items, floorMap[option.Floor])
//
//		newFloorMap := MakeNewFloorMap(floorMap, newFloor, oldFloor)
//		newFloorHash := FloorMapHash(newFloorMap)
//		if _, ok := floorCombinationHash[newFloorHash]; ok {
//			//fmt.Printf("Already seen %v, skipping.\n", newFloorHash)
//			continue
//		} else {
//			floorCombinationHash[newFloorHash] = true
//		}
//		newElevator := &RTGElevator{
//			Floor: option.GotoFloor,
//			Generators: option.Generators(),
//			Microchips: option.Microchips(),
//		}
//		fmt.Printf("New Elevator choice: %v with map %v\r", newElevator, newFloorMap)
//
//		nextBest := MoveRTGItems(newFloorMap, newElevator, moves)
//		if nextBest < best {
//			best = nextBest
//		}
//	}
//	return best
//}

func MakeNewFloorMap(floorMap map[int]*RTGFloor, newFloor *RTGFloor, oldFloor *RTGFloor) map[int]*RTGFloor {
	newFloorMap := map[int]*RTGFloor{}
	for i := 1; i <= maxFloor; i++ {
		if newFloor.Number == i {
			newFloorMap[i] = newFloor
		} else if oldFloor.Number == i {
			newFloorMap[i] = oldFloor
		} else {
			newFloorMap[i] = &RTGFloor{
				Number:     floorMap[i].Number,
				Generators: floorMap[i].Generators,
				Microchips: floorMap[i].Microchips,
			}
		}
	}
	return newFloorMap
}

func TopFloorHasEverything(floor *RTGFloor) bool {
	return floor.Number == maxFloor &&
		len(floor.Generators) == maxGenerators &&
		len(floor.Microchips) == maxMicrochips
}

type RTGOptions struct {
	Floor     int
	GotoFloor int
	Items     []string
}

func (o *RTGOptions) String() string {
	return fmt.Sprintf("Option %v->%v taking %v\n", o.Floor, o.GotoFloor, o.Items)
}

func (o *RTGOptions) Generators() []string {
	newGenerators := []string{}
	for _, item := range o.Items {
		if item[1] == 'G' {
			newGenerators = append(newGenerators, item)
		}
	}
	sort.Strings(newGenerators)
	return newGenerators
}

func (o *RTGOptions) Microchips() []string {
	newMicrochips := []string{}
	for _, item := range o.Items {
		if item[1] == 'M' {
			newMicrochips = append(newMicrochips, item)
		}
	}
	sort.Strings(newMicrochips)
	return newMicrochips
}

func GetOptionsForElevatorAndFloor(elevator *RTGElevator, floorMap map[int]*RTGFloor) []*RTGOptions {

	options := []*RTGOptions{}
	floor := floorMap[elevator.Floor]
	items := floor.Items()
	nextFloors := GetNextFloors(elevator.Floor)

	//fmt.Printf("Have %v Items to move\n", items)
	for _, f := range nextFloors {
		//fmt.Printf("Going to Floor %v\n", f)
		//Pick one item:
		for _, item := range items {
			//fmt.Printf("Taking %v\n", item)
			option := &RTGOptions{
				Floor:     elevator.Floor,
				GotoFloor: f,
				Items:     []string{item},
			}
			options = append(options, option)
		}

		//Pick two items:
		for i := 0; i < len(items); i++ {
			for j := i + 1; j < len(items); j++ {
				if i == j {
					continue
				}
				//fmt.Printf("Taking %v and %v\n", items[i], items[j])
				newItems := []string{items[i], items[j]}
				sort.Strings(newItems)
				option := &RTGOptions{
					Floor:     elevator.Floor,
					GotoFloor: f,
					Items:     newItems,
				}
				options = append(options, option)
			}
		}
	}

	return options
}

//func FriedOption(option *RTGOptions, floorMap map[int]*RTGFloor) bool {
//
//	if len(option.Items) == 0 {
//		return false
//	}
//
//	for _, item := range option.Items {
//		if LeavesDanglingFriedMicrochip(item, floorMap[option.Floor]) ||
//		   //ElevatorFriesMicrochip(item, option.Items) ||
//		   FriesMicrochipOnFloor(item, floorMap[option.GotoFloor]) {
//			return false
//		}
//	}
//	return true
//
//}

func NewFloorWithItems(items []string, floor *RTGFloor) (*RTGFloor, bool) {

	newGenerators := []string{}
	for _, g := range floor.Generators {
		newGenerators = append(newGenerators, g)
	}
	newMicrochips := []string{}
	for _, mc := range floor.Microchips {
		newMicrochips = append(newMicrochips, mc)
	}
	for _, item := range items {
		if item[1] == 'G' {
			newGenerators = append(newGenerators, item)
		} else {
			newMicrochips = append(newMicrochips, item)
		}
	}
	sort.Strings(newGenerators)
	sort.Strings(newMicrochips)

	newFloor := &RTGFloor{
		Number:     floor.Number,
		Generators: newGenerators,
		Microchips: newMicrochips,
	}

	return newFloor, newFloor.IsValid()

}

func (f *RTGFloor) IsValid() bool {
	// if generator is open, there must not be open microchips
	openMicrochip := false
	openGenerator := false

	for _, mc := range f.Microchips {
		nextMCOpen := true
		for _, g := range f.Generators {
			if mc[0] == g[0] {
				nextMCOpen = false // Found pair for MC
			}
		}
		openMicrochip = openMicrochip || nextMCOpen
	}

	for _, g := range f.Generators {
		nextGenOpen := true
		for _, mc := range f.Microchips {

			if mc[0] == g[0] {
				nextGenOpen = false //Found pair for MC
			}
		}
		openGenerator = openGenerator || nextGenOpen

	}

	if openGenerator && openMicrochip {
		//fmt.Printf("Invalid Floor: %v - %v\n", f.Generators, f.Microchips)
		return false
	}
	return true
}

func NewFloorWithoutItems(items []string, floor *RTGFloor) (*RTGFloor, bool) {

	newGenerators := []string{}
GenLoop:
	for _, g := range floor.Generators {
		for _, item := range items {
			if item == g {
				continue GenLoop //Removing
			}
		}
		newGenerators = append(newGenerators, g)
	}
	newMicrochips := []string{}
MCLoop:
	for _, mc := range floor.Microchips {
		for _, item := range items {
			if item == mc {
				continue MCLoop // Removing
			}
		}
		newMicrochips = append(newMicrochips, mc)
	}
	sort.Strings(newGenerators)
	sort.Strings(newMicrochips)

	newFloor := &RTGFloor{
		Number:     floor.Number,
		Generators: newGenerators,
		Microchips: newMicrochips,
	}

	return newFloor, newFloor.IsValid()

}

//
//func FriesMicrochipOnFloor(item string, floor *RTGFloor) bool {
//	fmt.Printf("Can I take %v to %v with %v?\n", item, floor.Number, floor.Generators)
//
//	if item[1] == 'G' {
//		//Taking Generator to new floor..
//		openMicrochip := false
//		for _, mc := range floor.Microchips {
//			if item[0] == mc[0] {
//				return false //Paired MC with G, we're ok
//			}
//			nextMCOpen := true
//			for _, g := range floor.Generators {
//				if mc[0] == g[0] {
//					nextMCOpen = false //This MC has a paired G
//				}
//			}
//			openMicrochip = openMicrochip || nextMCOpen
//		}
//		if openMicrochip {
//			return true // This G will fry an open MC
//		}
//
//	} else { // Bringing a MC to a new floor
//		openGenerator := false
//		for _, g := range floor.Generators {
//			if item[0] == g[0] {
//				return false // Paired MC with G, we're ok
//			}
//			nextOpenGenerator := true
//			for _, mc := range floor.Microchips {
//				if mc[0] == g[0] {
//					nextOpenGenerator = false // This G has a paired MC
//				}
//			}
//			openGenerator = openGenerator || nextOpenGenerator
//		}
//		if openGenerator {
//			return true // This MC will be fried by an open G
//		}
//	}
//	return false // OK. No open Generators to fry this MC, or no dangling MCs to be fried by this G.
//}

//func ElevatorFriesMicrochip(item string, items []string) bool {
//	if item[1] == 'G' {
//		return false
//	}
//	for _, other := range items {
//		if other == item {
//			continue
//		}
//		if other[1] == 'M' {
//			continue
//		}
//		if item[0] != other[0] {
//			return true
//		}
//	}
//	return false
//}
//
//func LeavesDanglingFriedMicrochip(item string, floor *RTGFloor) bool {
//	fmt.Printf("Can I remove %v from %v with %v and %v?\n", item, floor.Number, floor.Generators, floor.Microchips)
//
//	if len(floor.Generators) == 0 {
//		// No generators, no fried MCs
//		return false
//	}
//
//	//item is being removed.
//	if len(floor.Microchips) == 1 {
//		// that's us
//		return false
//	}
//
//	// if generator is open, there must not be open microchips
//	openMicrochip := false
//	openGenerator := false
//
//	if item[1] == 'M' {
//		// Taking microchip...
//		for _, mc := range floor.Microchips {
//			if mc == item {
//				continue //Don't check ourselves, we aren't staying
//			}
//			nextMCOpen := true
//			for _, g := range floor.Generators {
//				if g[0] == item[0] {
//					// item is being taken, leaving generator open
//					openGenerator = true
//				} else if mc[0] == g[0] {
//					nextMCOpen = false // Still found pair for MC
//				}
//			}
//			openMicrochip = openMicrochip || nextMCOpen
//		}
//	} else {
//		//Taking generator
//		for _, mc := range floor.Microchips {
//			nextMCOpen := true
//			for _, g := range floor.Generators {
//				if g == item {
//					continue // We're taking this generator
//				}
//				if mc[0] == g[0] {
//					nextMCOpen = false //Still found pair for MC
//				}
//			}
//			openMicrochip = openMicrochip || nextMCOpen
//		}
//	}
//
//	if openGenerator && openMicrochip {
//		fmt.Printf("Taking %v leaves dangling MC. %v - %v\n", item, floor.Generators, floor.Microchips)
//		return true
//	}
//	return false
//
//}

func GetNextFloors(floor int) []int {
	var nextFloors []int
	if floor == 1 {
		nextFloors = []int{2}
	} else if floor == 4 {
		nextFloors = []int{3}
	} else {
		nextFloors = []int{floor - 1, floor + 1}
	}

	return nextFloors
}
