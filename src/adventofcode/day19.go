package adventofcode

import "fmt"

func Day19() {
	if true {
		Day19Part2Slow()
		return
	}

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

func Day19Part2Slow() {
	numElves := 3017957
	//numElves := 5

	elves := make([]int, numElves)
	for i := 0; i < numElves; i++ {
		elves[i] = i + 1
	}

	fmt.Printf("Num elves: %v\n", len(elves))
	for len(elves) > 1 {
		victimPtr := len(elves) / 2
		fmt.Printf("(%v), %v takes %v\r", len(elves), elves[0], elves[victimPtr])
		newElves := elves[1:victimPtr]
		newElves = append(newElves, elves[victimPtr+1:]...)
		newElves = append(newElves, elves[0])
		elves = newElves
		//fmt.Printf("New: %v\n", elves)
		if len(elves) == 2 {
			elves = []int{elves[0]}
		}
	}

	fmt.Printf("\nLast Elf: %v\n", elves[0])
}

func Day19Part2Slower() {
	numElves := 3017957
	//numElves := 5

	elves := make([]bool, numElves)
	for i := 0; i < numElves; i++ {
		elves[i] = true
	}
	ptr := 0
	acrossPtr := numElves / 2
	numElvesAlive := numElves

	for numElvesAlive > 2 {

		elves[acrossPtr] = false
		numElvesAlive -= 1
		//if numElvesAlive % 1000 == 0 {
		//	fmt.Printf("(%v)\tElf %v takes out %v\r", numElvesAlive, ptr+1, acrossPtr+1)
		//}
		ptr = (ptr + 1) % numElves
		//fmt.Printf("Next elf may be %v\n", ptr+1)

		//if true {return}
		for !elves[ptr] {
			ptr = (ptr + 1) % numElves
			//fmt.Printf("Next elf may really be %v\n", ptr+1)
			acrossPtr = (acrossPtr + 1) % numElves
		}
		//fmt.Printf("Elves remaining: %v\n", numElvesAlive)

		//Find next 'across':
		numRemainingCount := 0
		acrossPtr = ptr + 1
		for numRemainingCount < (numElvesAlive / 2) {
			//fmt.Printf("Moving victim ptr to %v\n", acrossPtr+1)
			acrossPtr = (acrossPtr + 1) % numElves
			if elves[acrossPtr] {
				numRemainingCount++
			}
			//return

		}
		for !elves[acrossPtr] {
			acrossPtr = (acrossPtr + 1) % numElves
			//fmt.Printf("Next victim may be %v\n", acrossPtr+1)
		}
		//Counted halfway around from ptr.
		//fmt.Printf("Next victim is %v\n", acrossPtr+1)
	}

	fmt.Printf("\nLast Elf: %v\n", ptr+1)
}
