package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

func Day12() {
	day := "12"
	filename := fmt.Sprintf("data/day%vinput.txt", day)
	input := readFileAsLines(filename)

	instructions := []*ABInstr{}
	for _, line := range input {
		fmt.Printf("Line: %v\n", line)
		instructions = append(instructions, NewABInstr(line))
	}

	comp := &ABComputer{
		Registers: map[string]int{
			//"a": 0, "b": 0, "c": 0, "d": 0, // Part 1
			"a": 0, "b": 0, "c": 1, "d": 0, // Part 2
		},
	}

	RunABComputer(comp, instructions)

	fmt.Printf("Result: %v\n", comp.Registers["a"])
}

func RunABComputer(comp *ABComputer, instructions []*ABInstr) {

	ptr := 0
	for ptr < len(instructions) {
		ptr = ProcessInstruction(comp, instructions[ptr], ptr)
	}
}

func ProcessInstruction(comp *ABComputer, instr *ABInstr, ptr int) int {
	fmt.Printf("[%v] Instr: %v, Comp: %v       \r", ptr, instr, comp)
	switch instr.Name {
	case "cpy":
		val := GetValueForInstr(instr.XVal, comp)
		comp.Registers[instr.YVal] = val
		ptr += 1
	case "jnz":
		val := GetValueForInstr(instr.XVal, comp)
		if val != 0 {
			offset, _ := strconv.Atoi(instr.YVal)
			ptr += offset
		} else {
			ptr += 1
		}

	case "inc":
		comp.Registers[instr.XVal] += 1
		ptr += 1

	case "dec":
		comp.Registers[instr.XVal] -= 1
		ptr += 1
	}
	return ptr
}

func GetValueForInstr(label string, comp *ABComputer) int {
	if label == "a" || label == "b" || label == "c" || label == "d" {
		return comp.Registers[label]
	} else {
		num, _ := strconv.Atoi(label)
		return num
	}
}

type ABComputer struct {
	Registers map[string]int
}

func (c *ABComputer) String() string {
	return fmt.Sprintf("%v", c.Registers)
}

type ABInstr struct {
	Name string
	XVal string
	YVal string
}

func (i *ABInstr) String() string {
	return fmt.Sprintf("%v %v %v", i.Name, i.XVal, i.YVal)
}

func NewABInstr(line string) *ABInstr {
	tokens := strings.Split(line, " ")
	switch tokens[0] {
	case "cpy", "jnz":
		return &ABInstr{
			Name: tokens[0],
			XVal: tokens[1],
			YVal: tokens[2],
		}
	case "inc", "dec":
		return &ABInstr{
			Name: tokens[0],
			XVal: tokens[1],
		}
	default:
		panic("Unhandled instruction")
	}
}
