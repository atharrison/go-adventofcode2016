package adventofcode

import (
	"fmt"
)

func Day16() {

	// Part 2
	if true {
		Day16Bytes()
		return
	}

	input := "10010000000110000"

	// Part 1
	diskSize := 272

	// Part 2
	//diskSize := 35651584
	// Don't do Part 2 with strings.

	for len(input) < diskSize {
		input = dragonCurveLengthen(input)
		fmt.Printf("Lengthen to %v\r", len(input))
	}

	cropped := input[0:diskSize]
	fmt.Printf("\nCropping %v to %v\n", len(cropped), diskSize)

	//fmt.Printf("Cropped: (%v) = %v\n", cropped, len(cropped))

	result := reduceDragon(cropped)
	fmt.Println(result)

}

func reduceDragon(input string) string {
	output := ""
	for len(input)%2 == 0 {
		fmt.Printf("Reduce from %v\n", input)
		//fmt.Printf("Reducing from %v\r", len(input))
		output = ""
		for i := 0; i < len(input)-1; i += 2 {
			if input[i] == input[i+1] {
				output = output + "1"
			} else {
				output = output + "0"
			}
		}
		input = output
	}
	return output
}

func dragonCurveLengthen(input string) string {
	a := input
	b := SwapDigits(Reverse(input))

	return fmt.Sprintf("%v0%v", a, b)
}

func SwapDigits(input string) string {
	output := ""
	for _, c := range input {
		if c == '1' {
			output = output + "1"
		} else {
			output = output + "0"
		}
	}
	return output
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//=============
// Part 2, Using byte arrays, not strings.

func Day16Bytes() {

	input := []byte{
		1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0,
	}

	// Part 1
	//diskSize := 272

	// Part 2
	diskSize := 35651584

	for len(input) < diskSize {
		input = dragonCurveLengthenBytes(input)
		fmt.Printf("Lengthen to %v\r", len(input))
	}

	cropped := input[0:diskSize]
	fmt.Printf("\nCropping %v to %v\n", len(input), diskSize)

	//fmt.Printf("Cropped: (%v) = %v\n", cropped, len(cropped))

	result := reduceDragonBytes(cropped)
	for _, d := range result {
		fmt.Printf("%v", d)
	}
	fmt.Println()
	//fmt.Println(result)

}

func reduceDragonBytes(input []byte) []byte {
	var output []byte
	//fmt.Printf("Reduce from %v\n", input)
	for len(input)%2 == 0 {
		fmt.Printf("Reducing from %v\r", len(input))
		if len(input) < 50 {
			fmt.Printf("Reduce from %v\n", input)
		}
		size := len(input) / 2
		output = make([]byte, size)

		for i := 0; i < len(input)-1; i += 2 {
			idx := i / 2
			if input[i] == input[i+1] {
				output[idx] = 1
			} else {
				output[idx] = 0
			}
		}
		input = output
		//fmt.Printf("Reduce from %v\n", input)
	}
	return output
}

func dragonCurveLengthenBytes(input []byte) []byte {
	a := input
	b := SwapDigitsBytes(ReverseBytes(input))

	output := a
	output = append(output, 0)
	for _, c := range b {
		output = append(output, c)
	}

	return output
}

func SwapDigitsBytes(input []byte) []byte {
	size := len(input)
	output := make([]byte, size)
	for i := 0; i < size; i++ {
		if input[i] == 0 {
			output[i] = 1
		}
		//already 0
		//} else {
		//	output[i] = 0
		//}
	}
	return output
}

func ReverseBytes(arr []byte) []byte {

	size := len(arr)
	output := make([]byte, size)

	for i := 0; i < size; i++ {
		output[size-i-1] = arr[i]
	}
	//fmt.Printf("Reverse %v to %v\n", arr, output)
	return output
}
