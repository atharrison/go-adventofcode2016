package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
	//"math"
)

var day8Grid [][]int

var xSize = 50
var ySize = 6

//var xSize = 7
//var ySize = 3

func Day08() {
	day := "08"
	filename := fmt.Sprintf("data/day%v_bonus.txt", day)
	//filename := fmt.Sprintf("data/day%vinput.txt", day)
	//filename := fmt.Sprintf("data/day%vinput_sample.txt", day)
	input := readFileAsLines(filename)

	day8Grid = make([][]int, ySize)
	for y := 0; y < ySize; y++ {
		day8Grid[y] = make([]int, xSize)
	}

	var instructions []*Day08Instruction
	for _, line := range input {
		fmt.Printf("Line: %v\n", line)
		instr := NewDay08Instruction(line)
		instructions = append(instructions, instr)
		processDay08Instr(instr)
		printDay08Grid(day8Grid)
	}

	//fmt.Printf("%v\n", instructions)
	//for i := 0; i < xSize; i++ {
	//	fmt.Println(day8Grid[i])
	//}
	//printDay08Grid(day8Grid)
	total := countDay08Grid(day8Grid)
	fmt.Printf("%v\n", total)
}

type Day08Instruction struct {
	Type string
	A    int
	B    int
}

func printDay08Grid(grid [][]int) {
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {

			// Part 2, make it easier to see:
			if day8Grid[y][x] == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}

			//fmt.Printf("%v", day8Grid[y][x])
		}
		fmt.Println()
	}
}
func processDay08Instr(instr *Day08Instruction) {
	switch instr.Type {
	case "rect":
		for x := 0; x < instr.A; x++ {
			for y := 0; y < instr.B; y++ {
				day8Grid[y][x] = 1
			}
		}
	case "row":
		moveRow(instr.A, instr.B)
	case "column":
		moveColumn(instr.A, instr.B)
	}
}

func moveRow(y int, delta int) {
	prevRow := make([]int, xSize)
	for r := 0; r < xSize; r++ {
		prevRow[r] = day8Grid[y][r]
	}
	for i := 0; i < xSize; i++ {
		day8Grid[y][(i+delta)%xSize] = prevRow[i]
	}
}

func moveColumn(x int, delta int) {

	prevCol := make([]int, ySize)
	for c := 0; c < ySize; c++ {
		prevCol[c] = day8Grid[c][x]
	}
	for i := 0; i < ySize; i++ {
		day8Grid[(i+delta)%ySize][x] = prevCol[i]
	}
}

func countDay08Grid(grid [][]int) int {
	total := 0
	for i := 0; i < xSize; i++ {
		for j := 0; j < ySize; j++ {
			total += day8Grid[j][i]
		}
	}
	return total
}
func (i *Day08Instruction) String() string {
	return fmt.Sprintf("[%v %v %v]", i.Type, i.A, i.B)
}

func NewDay08Instruction(line string) *Day08Instruction {
	tokens := strings.Split(line, " ")
	fmt.Println(tokens)
	if tokens[0] == "rect" {

		nums := strings.Split(tokens[1], "x")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		fmt.Println("Returning rect Instr")
		return &Day08Instruction{
			Type: "rect",
			A:    a,
			B:    b,
		}
	}
	//} else if strings.HasPrefix("rotate", line) {

	if tokens[1] == "column" {
		a, _ := strconv.Atoi(tokens[2][2:])
		b, _ := strconv.Atoi(tokens[4])
		return &Day08Instruction{
			Type: "column",
			A:    a,
			B:    b,
		}
	}
	//else if tokens[1] == "row" {

	a, _ := strconv.Atoi(tokens[2][2:])
	b, _ := strconv.Atoi(tokens[4])
	return &Day08Instruction{
		Type: "row",
		A:    a,
		B:    b,
	}
}
