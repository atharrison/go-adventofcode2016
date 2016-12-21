package adventofcode

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func Day21() {
	day := "21"
	filename := fmt.Sprintf("data/day%vinput.txt", day)
	//filename := fmt.Sprintf("data/day%vinput_sample.txt", day)
	input := readFileAsLines(filename)

	// Part 1
	//toScramble := "abcdefgh"
	//toScramble := "abcde"

	//instructions := []*Day21Instruction{}
	//for _, line := range input {
	//	instr := ParseDay21Instruction(line)
	//	instructions = append(instructions, instr)
	//	fmt.Printf("%v\n", instr)
	//}

	// Part 2
	instructions := []*Day21Instruction{}
	toScramble := "fbgdceah"
	//toScramble := "decab"
	for i := len(input) - 1; i >= 0; i-- {
		line := input[i]
		instr := ReverseDay21Instruction(ParseDay21Instruction(line))
		instructions = append(instructions, instr)
		fmt.Printf("%v\n", instr)
	}

	result := ProcessDay21Instructions(toScramble, instructions)
	fmt.Printf("%v\n", result)
}

func ReverseDay21Instruction(instr *Day21Instruction) *Day21Instruction {
	switch instr.Action {
	case "swapP":
		return instr
	case "swapL":
		return instr
	case "rotateL":
		return &Day21Instruction{
			Action: "rotateR",
			PosX:   instr.PosX,
		}
	case "rotateR":
		return &Day21Instruction{
			Action: "rotateL",
			PosX:   instr.PosX,
		}
	case "rotateB":
		return &Day21Instruction{
			Action:  "rotateBBack",
			LetterX: instr.LetterX,
		}
	case "reverse":
		return instr
	case "move":
		return &Day21Instruction{
			Action: instr.Action,
			PosX:   instr.PosY,
			PosY:   instr.PosX,
		}
	default:
		fmt.Printf("Unknown instr %v\n", instr.Action)
		return nil
	}
}

func ProcessDay21Instructions(toScramble string, instructions []*Day21Instruction) string {
	letters := strings.Split(toScramble, "")
	fmt.Printf("Letters: %v\n", letters)
	for _, instr := range instructions {
		fmt.Printf("Letters now: %v, applying %v\n", letters, instr)
		switch instr.Action {
		case "swapP":
			temp := letters[instr.PosX]
			letters[instr.PosX] = letters[instr.PosY]
			letters[instr.PosY] = temp
		case "swapL":
			posX := FindPositionForLetter(letters, instr.LetterX)
			posY := FindPositionForLetter(letters, instr.LetterY)
			temp := letters[posX]
			letters[posX] = letters[posY]
			letters[posY] = temp
		case "rotateL":
			for i := 0; i < instr.PosX; i++ {
				letters = RotateLeft(letters)
			}
		case "rotateR":
			for i := 0; i < instr.PosX; i++ {
				letters = RotateRight(letters)
			}
		case "rotateB":
			/**
			rotate based on position of letter X means that the whole string should be rotated to the
			right based on the index of letter X (counting from 0) as determined before this
			instruction does any rotations. Once the index is determined, rotate the string
			to the right one time, plus a number of times equal to that index,
			plus one additional time if the index was at least 4.
			*/
			letters, _ = RotateB(letters, instr.LetterX)
		case "rotateBBack":
			// Rotate left until executing a 'rotateB' would result in the original
			fmt.Printf("RotateBBack start with %v\n", letters)
			checkLetters := RotateLeft(letters)
			rotated, rotatedRightCount := RotateB(letters, instr.LetterX)

			for !reflect.DeepEqual(letters, rotated) {
				checkLetters = RotateLeft(checkLetters)
				rotated, rotatedRightCount = RotateB(checkLetters, instr.LetterX)
				fmt.Printf("Rotating left produces %v, which RotateB would result in %v (%v)\n", checkLetters, rotated, rotatedRightCount)
			}
			letters = checkLetters
		case "reverse":
			letters = ReverseLetters(letters, instr.PosX, instr.PosY)
		case "move":
			/**
			move position X to position Y means that the letter
			which is at index X should be removed from the string,
			then inserted such that it ends up at index Y
			*/
			letters = MoveLetter(letters, instr.PosX, instr.PosY)
		}
	}

	result := ""
	for i := 0; i < len(letters); i++ {
		result += letters[i]
	}
	return result
}

func RotateB(letters []string, letter string) ([]string, int) {
	posX := FindPositionForLetter(letters, letter)
	result := letters
	rotateRightCount := posX + 1
	if posX >= 4 {
		rotateRightCount++
	}
	for i := 0; i < rotateRightCount; i++ {
		result = RotateRight(result)
	}
	return result, rotateRightCount
}

func MoveLetter(letters []string, start int, end int) []string {
	fmt.Printf("Move starting with %v, moving %v to %v\n", letters, start, end)
	toMove := letters[start]
	tempLetters := make([]string, len(letters)-1)
	if start == len(letters)-1 {
		for i := 0; i < start; i++ {
			tempLetters[i] = letters[i]
		}
	} else {
		for i := 0; i <= start; i++ {
			tempLetters[i] = letters[i]
		}
		for i := start + 1; i < len(letters); i++ {
			tempLetters[i-1] = letters[i]
		}
	}

	//tempLetters = append(tempLetters, letters[start+1:]...)
	fmt.Printf("TempLetters, sans %v: %v\n", toMove, tempLetters)

	result := make([]string, len(letters))
	for i := 0; i < end; i++ {
		result[i] = tempLetters[i]
	}
	result[end] = toMove
	for i := end; i < len(tempLetters); i++ {
		result[i+1] = tempLetters[i]
	}

	fmt.Printf("Moving %v from %v to %v, %v to %v\n", toMove, start, end, letters, result)
	return result
}

func ReverseStringSlice(arr []string) []string {

	size := len(arr)
	output := make([]string, size)

	for i := 0; i < size; i++ {
		output[size-i-1] = arr[i]
	}
	fmt.Printf("Reverse %v to %v\n", arr, output)
	return output
}

func ReverseLetters(letters []string, start int, end int) []string {

	result := make([]string, len(letters))
	fmt.Printf("Reversing %v (%v-%v)\n", letters, start, end)
	for i := 0; i < start; i++ {
		result[i] = letters[i]
	}
	reversed := ReverseStringSlice(letters[start : end+1])
	i := start
	for _, v := range reversed {
		result[i] = v
		i++
	}
	for i := end + 1; i < len(letters); i++ {
		result[i] = letters[i]
	}

	fmt.Printf("Reversed %v (%v-%v)to %v\n", letters, start, end, result)
	return result
}

func FindPositionForLetter(letters []string, letter string) int {
	for idx, l := range letters {
		if l == letter {
			return idx
		}
	}
	fmt.Printf("WTF? Can't find letter %v in %v\n", letter, letters)
	os.Exit(1)
	return 0 // Won't get here
}

func RotateLeft(letters []string) []string {
	tmp := letters[0]
	result := letters[1:]
	result = append(result, tmp)
	//fmt.Printf("Rotate Left turns %v into %v\n", letters, result)
	return result
}

func RotateRight(letters []string) []string {
	result := []string{letters[len(letters)-1]}
	result = append(result, letters[0:len(letters)-1]...)
	//fmt.Printf("Rotate Right turns %v into %v\n", letters, result)
	return result
}

type Day21Instruction struct {
	InstrStr string
	Action   string
	LetterX  string
	LetterY  string
	PosX     int
	PosY     int
}

func (d *Day21Instruction) String() string {
	return fmt.Sprintf("%v: %v %v %v %v", d.Action, d.LetterX, d.LetterY, d.PosX, d.PosY)
}

func ParseDay21Instruction(line string) *Day21Instruction {
	// swap position X with position Y
	// swap letter X with letter Y
	// rotate left/right X steps
	// rotate based on position of letter X
	// reverse positions X through Y
	// move position X to position Y
	tokens := strings.Split(line, " ")

	if tokens[0] == "swap" {
		if tokens[1] == "position" {
			posX, _ := strconv.Atoi(tokens[2])
			posY, _ := strconv.Atoi(tokens[5])
			return &Day21Instruction{
				InstrStr: line,
				Action:   "swapP",
				PosX:     posX,
				PosY:     posY,
			}
		} else {
			return &Day21Instruction{
				InstrStr: line,
				Action:   "swapL",
				LetterX:  tokens[2],
				LetterY:  tokens[5],
			}

		}
	}
	if tokens[0] == "rotate" {
		if tokens[1] == "left" {
			posX, _ := strconv.Atoi(tokens[2])
			return &Day21Instruction{
				InstrStr: line,
				Action:   "rotateL",
				PosX:     posX,
			}
		} else if tokens[1] == "right" {
			posX, _ := strconv.Atoi(tokens[2])
			return &Day21Instruction{
				InstrStr: line,
				Action:   "rotateR",
				PosX:     posX,
			}
		} else {
			return &Day21Instruction{
				InstrStr: line,
				Action:   "rotateB",
				LetterX:  tokens[6],
			}
		}
	}
	if tokens[0] == "reverse" {
		posX, _ := strconv.Atoi(tokens[2])
		posY, _ := strconv.Atoi(tokens[4])
		return &Day21Instruction{
			InstrStr: line,
			Action:   "reverse",
			PosX:     posX,
			PosY:     posY,
		}
	}
	if tokens[0] == "move" {
		posX, _ := strconv.Atoi(tokens[2])
		posY, _ := strconv.Atoi(tokens[5])
		return &Day21Instruction{
			InstrStr: line,
			Action:   "move",
			PosX:     posX,
			PosY:     posY,
		}
	}

	return &Day21Instruction{
		InstrStr: line,
	}
}
