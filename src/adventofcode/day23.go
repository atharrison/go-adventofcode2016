package adventofcode

// Moved to Day 25
//
//import (
//	"fmt"
//	"strconv"
//	"strings"
//	//"os"
//	"os"
//)
//
//func Day23() {
//	day := "23"
//	//filename := fmt.Sprintf("data/day%vinput_sample2.txt", day)
//	//filename := fmt.Sprintf("data/day%vinput_sample.txt", day)
//	//filename := fmt.Sprintf("data/day%vinput.txt", day)
//	filename := fmt.Sprintf("data/day%vinput_handoptimized2.txt", day)
//	input := readFileAsLines(filename)
//
//	instructions := []*ABInstr{}
//	for _, line := range input {
//		fmt.Printf("Line: %v\n", line)
//		instructions = append(instructions, NewABInstr(line))
//	}
//
//	//instructions = OptimizeInstructions(instructions)
//	//for idx, instr := range instructions {
//	//	fmt.Printf("%v: %v\n", idx, instr)
//	//}
//
//	comp := &ABComputer{
//		Registers: map[string]int{
//			//"a": 7, "b": 0, "c": 0, "d": 0, // Part 1
//			"a": 12, "b": 0, "c": 1, "d": 0, // Part 2
//		},
//	}
//
//	RunABComputer(comp, instructions)
//
//	fmt.Printf("\n\nResult: %v\n", comp.Registers["a"])
//}
//
//func RunABComputer(comp *ABComputer, instructions []*ABInstr) {
//
//	ptr := 0
//	for ptr >= 0 && ptr < len(instructions) {
//		ptr = ProcessInstruction(comp, instructions[ptr], ptr, instructions)
//	}
//}
//
//func ProcessInstruction(comp *ABComputer, instr *ABInstr, ptr int, instructions []*ABInstr) int {
//	//Ha! Best optimization is simply not to print every line of execution
//	fmt.Printf("[%v] Instr: %v, Comp: %v       \n", ptr, instr, comp)
//	switch instr.Name {
//	case "cpy":
//		val := GetValueForInstr(instr.XVal, comp)
//		comp.Registers[instr.YVal] = val
//		ptr += 1
//	case "jnz":
//		val := GetValueForInstr(instr.XVal, comp)
//		if val != 0 {
//			//offset, _ := strconv.Atoi(instr.YVal)
//			offset := GetValueForInstr(instr.YVal, comp)
//			//fmt.Printf("Jumpting to %v (offset: %v)\n", offset+ptr, offset)
//			ptr += offset
//		} else {
//			ptr += 1
//		}
//
//	case "inc":
//		comp.Registers[instr.XVal] += 1
//		ptr += 1
//
//	case "dec":
//		comp.Registers[instr.XVal] -= 1
//		ptr += 1
//	case "tgl":
//		val := GetValueForInstr(instr.XVal, comp)
//		ToggleInstruction(ptr+val, instructions)
//		ptr += 1
//
//	// New optimized instructions:
//	case "incv":
//		val := GetValueForInstr(instr.YVal, comp)
//		comp.Registers[instr.XVal] += val
//		ptr += 1
//	case "decv":
//		val := GetValueForInstr(instr.YVal, comp)
//		comp.Registers[instr.XVal] -= val
//		ptr += 1
//	case "mtp":
//		// Multiply YVal and ZVal, and add it to XVal
//		oper1 := GetValueForInstr(instr.YVal, comp)
//		oper2 := GetValueForInstr(instr.ZVal, comp)
//		comp.Registers[instr.XVal] += oper1 * oper2
//		ptr += 1
//	case "noop":
//		//Do nothing
//		ptr += 1
//	}
//
//	return ptr
//}
//
//func ToggleInstruction(ptr int, instructions []*ABInstr) {
//	if ptr >= len(instructions) {
//		fmt.Printf("Toggle NOOP, beyond instr tape\n")
//		return
//	}
//	instrToToggle := instructions[ptr]
//	switch instrToToggle.Name {
//	case "inc":
//		//One-arg instr
//		instructions[ptr] = &ABInstr{
//			Name: "dec",
//			XVal: instrToToggle.XVal,
//		}
//	case "dec", "incv", "decv", "tgl":
//		//One-arg instr
//		instructions[ptr] = &ABInstr{
//			Name: "inc",
//			XVal: instrToToggle.XVal,
//		}
//
//	case "jnz":
//		//Two-arg instr
//		instructions[ptr] = &ABInstr{
//			Name: "cpy",
//			XVal: instrToToggle.XVal,
//			YVal: instrToToggle.YVal,
//		}
//	case "cpy":
//		//Two-arg instr
//
//		instructions[ptr] = &ABInstr{
//			Name: "jnz",
//			XVal: instrToToggle.XVal,
//			YVal: instrToToggle.YVal,
//		}
//	case "mtp":
//		// Doh, toggle bites us...
//		fmt.Printf("Toggle trying to toggle %v", instructions[ptr])
//		os.Exit(1)
//	case "noop":
//		// NOTE: Since Toggle might want to modify an instr we optimized out,
//		// We need to insert NOOP type instructions in those places.
//		// But if toggle would try to toggle those, then things would get screwy.
//		fmt.Printf("Toggle trying to toggle %v", instructions[ptr])
//		os.Exit(1)
//	}
//
//	//fmt.Printf("Toggling %v to %v\n", instrToToggle, instructions[ptr])
//}
//
//func GetValueForInstr(label string, comp *ABComputer) int {
//	if label == "a" || label == "b" || label == "c" || label == "d" {
//		return comp.Registers[label]
//	} else {
//		num, _ := strconv.Atoi(label)
//		return num
//	}
//}
//
//type ABComputer struct {
//	Registers map[string]int
//}
//
//func (c *ABComputer) String() string {
//	return fmt.Sprintf("%v", c.Registers)
//}
//
//type ABInstr struct {
//	Name string
//	XVal string
//	YVal string
//	ZVal string
//}
//
//func (i *ABInstr) String() string {
//	return fmt.Sprintf("%v %v %v %v", i.Name, i.XVal, i.YVal, i.ZVal)
//}
//
//func NewABInstr(line string) *ABInstr {
//	tokens := strings.Split(line, " ")
//	switch tokens[0] {
//	case "mtp":
//		return &ABInstr{
//			Name: tokens[0],
//			XVal: tokens[1],
//			YVal: tokens[2],
//			ZVal: tokens[3],
//		}
//	case "cpy", "jnz":
//		return &ABInstr{
//			Name: tokens[0],
//			XVal: tokens[1],
//			YVal: tokens[2],
//		}
//	case "inc", "dec", "incv", "decv", "tgl":
//		return &ABInstr{
//			Name: tokens[0],
//			XVal: tokens[1],
//		}
//	case "noop":
//		return &ABInstr{
//			Name: tokens[0],
//		}
//	default:
//		panic("Unhandled instruction")
//	}
//}
