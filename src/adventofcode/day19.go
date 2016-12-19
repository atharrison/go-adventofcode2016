package adventofcode

import "fmt"

func Day19() {
	numElves := 3017957
	//numElves := 5

	numElvesWithPresents := 0
	presents := make([]int, numElves)
	for i := 0; i < numElves; i++ {
		presents[i] = 1
		numElvesWithPresents += 1
	}

	ptr := 0
	for numElvesWithPresents > 1 {
		//fmt.Printf("Checking Elf %v\n", ptr+1)
		if presents[ptr] > 0 {
			takePtr := (ptr + 1) % numElves
			for presents[takePtr] == 0 {
				takePtr = (takePtr + 1) % numElves
			}
			presents[ptr] += presents[takePtr]
			presents[takePtr] = 0
			fmt.Printf("Elf %v has %v\r", ptr+1, presents[ptr])
			ptr++
			numElvesWithPresents--
		} else {
			ptr++
		}
		ptr = ptr % numElves
	}

	fmt.Printf("\nLast Elf: %v\n", ptr)
}
