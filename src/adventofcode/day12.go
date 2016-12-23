package adventofcode

//
//import (
//	"fmt"
//	"strconv"
//	"strings"
//)
//
//func Day12() {
//	day := "12"
//	filename := fmt.Sprintf("data/day%vinput.txt", day)
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
//			//"a": 0, "b": 0, "c": 0, "d": 0, // Part 1
//			"a": 0, "b": 0, "c": 1, "d": 0, // Part 2
//		},
//	}
//
//	RunABComputer(comp, instructions)
//
//	fmt.Printf("\n\nResult: %v\n", comp.Registers["a"])
//}
//
//func OptimizeInstructions(instructions []*ABInstr) []*ABInstr {
//	optimized := []*ABInstr{}
//
//	for idx, instr := range instructions {
//		if idx > 1 {
//			if instr.Name == "jnz" && instr.YVal == "-2" &&
//				instr.XVal == instructions[idx-1].XVal &&
//				instructions[idx-1].Name == "dec" &&
//				instructions[idx-2].Name == "inc" {
//				optimized = optimized[0 : len(optimized)-2]
//				newInstr := &ABInstr{
//					Name: "incv",
//					XVal: instructions[idx-2].XVal,
//					YVal: instructions[idx-1].XVal,
//				}
//				optimized = append(optimized, newInstr)
//				newInstr = &ABInstr{
//					Name: "cpy",
//					XVal: "0",
//					YVal: instructions[idx-1].XVal,
//				}
//				optimized = append(optimized, newInstr)
//				//TODO: Adjust any jumps that would have been affected by this optimization.
//			} else {
//				optimized = append(optimized, instr)
//			}
//		} else {
//			optimized = append(optimized, instr)
//		}
//	}
//	return optimized
//}
//
//func RunABComputer(comp *ABComputer, instructions []*ABInstr) {
//
//	ptr := 0
//	for ptr < len(instructions) {
//		ptr = ProcessInstruction(comp, instructions[ptr], ptr)
//	}
//}
//
//func ProcessInstruction(comp *ABComputer, instr *ABInstr, ptr int) int {
//	//Ha! Best optimization is simply not to print every line of execution
//	//fmt.Printf("[%v] Instr: %v, Comp: %v       \r", ptr, instr, comp)
//	switch instr.Name {
//	case "cpy":
//		val := GetValueForInstr(instr.XVal, comp)
//		comp.Registers[instr.YVal] = val
//		ptr += 1
//	case "jnz":
//		val := GetValueForInstr(instr.XVal, comp)
//		if val != 0 {
//			offset, _ := strconv.Atoi(instr.YVal)
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
//	}
//	return ptr
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
//}
//
//func (i *ABInstr) String() string {
//	return fmt.Sprintf("%v %v %v", i.Name, i.XVal, i.YVal)
//}
//
//func NewABInstr(line string) *ABInstr {
//	tokens := strings.Split(line, " ")
//	switch tokens[0] {
//	case "cpy", "jnz":
//		return &ABInstr{
//			Name: tokens[0],
//			XVal: tokens[1],
//			YVal: tokens[2],
//		}
//	case "inc", "dec", "incv", "decv":
//		return &ABInstr{
//			Name: tokens[0],
//			XVal: tokens[1],
//		}
//	default:
//		panic("Unhandled instruction")
//	}
//}
