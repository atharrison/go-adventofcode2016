package adventofcode

import (
	"crypto/md5"
	"fmt"
)

var computedHashes = make(map[int]string)

func Day14() {
	salt := "ngcjuoqr"
	//salt := "abc"
	fmt.Printf("%v\n", salt)

	hashKeyCount := 0
	index := 0
	for hashKeyCount < 64 {
		toHash := fmt.Sprintf("%v%v", salt, index)
		nextHash := fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
		// Part 2- strech hashes
		strechedHash := strechedHash(index, nextHash, 2016)
		if tripleChar, ok := hashContainsTriple(strechedHash); ok {
			fmt.Printf("Index %v has a triple %v\r", index, string(tripleChar))
			good, idx := quintupleInNext1000(index, salt, tripleChar)
			if good {
				fmt.Printf("\nFound [%v] at idx %v from %v %v as %v from %v\n", string(tripleChar), idx, nextHash, hashKeyCount, index, toHash)
				hashKeyCount++
			}
		}
		index++
	}
	index -= 1
	fmt.Printf("Index: %v\n", index)
}

func strechedHash(index int, h string, num int) string {

	// Part 2's optimization key- don't recompute stretched hashes!
	if sh, ok := computedHashes[index]; ok {
		return sh
	}
	for i := 0; i < num; i++ {
		h = fmt.Sprintf("%x", md5.Sum([]byte(h)))
	}
	computedHashes[index] = h
	return h
}

func quintupleInNext1000(offset int, salt string, c byte) (bool, int) {
	for i := 1; i < 1001; i++ {

		toHash := fmt.Sprintf("%v%v", salt, offset+i)
		h := fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
		// Part 2, Also stretch hashes here:
		strechedHash := strechedHash(offset+i, h, 2016)
		if hashContainsQuintuple(strechedHash, c) {
			return true, offset + i
		}
	}
	return false, 0
}

func hashContainsTriple(h string) (byte, bool) {
	for i := 0; i < len(h)-2; i++ {
		c := h[i]
		if h[i+1] == c && h[i+2] == c {
			return c, true
		}
	}
	return 0, false
}

func hashContainsQuintuple(h string, c byte) bool {
	for i := 0; i < len(h)-4; i++ {
		if h[i] == c && h[i+1] == c && h[i+2] == c && h[i+3] == c && h[i+4] == c {
			return true
		}
	}
	return false
}
