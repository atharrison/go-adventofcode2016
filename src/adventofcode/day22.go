package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

func Day22() {
	day := "22"
	filename := fmt.Sprintf("data/day%vinput.txt", day)
	input := readFileAsLines(filename)

	firstLine := true

	maxX := 0
	maxY := 0
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
			maxX = gc.Ypos
		}
		fmt.Println(gc.String())
	}

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
	fmt.Printf("ValidNodeCound: %v\n", validNodeCount)
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

func (gc *GridComputer) String() string {
	return fmt.Sprintf("[x%v-y%v]\tU: %v\tA: %v", gc.Xpos, gc.Ypos, gc.Used, gc.Avail)
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
