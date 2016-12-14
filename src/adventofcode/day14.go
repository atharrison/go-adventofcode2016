package adventofcode

import (
	"crypto/md5"
	"fmt"
	"os"
)

func Day14() {
	salt := "ngcjuoqr"
	//salt := "abc"
	fmt.Printf("%v\n", salt)

	hashKeyCount := 0
	index := 0
	for hashKeyCount < 64 {
		toHash := fmt.Sprintf("%v%v", salt, index)
		nextHash := fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
		if tripleChar, ok := hashContainsTriple(nextHash); ok {
			//fmt.Printf("Index %v has a triple %v\n", index, string(tripleChar))
			if quintupleInNext1000(index, salt, tripleChar) {
				fmt.Printf("Found [%v] from %v %v as %v from %v\n", string(tripleChar), nextHash, hashKeyCount, index, toHash)
				hashKeyCount++
			}
		}
		index++
	}
	index -= 1
	fmt.Printf("Index: %v\n", index)
}

func quintupleInNext1000(offset int, salt string, c byte) bool {
	for i := 1; i < 1001; i++ {

		toHash := fmt.Sprintf("%v%v", salt, offset+i)
		h := fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
		if false && offset == 39 {
			toHash := fmt.Sprintf("%v%v", salt, 816)
			h := fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
			fmt.Printf("39: %v for index 816\n", h)
			os.Exit(0)
		}
		if hashContainsQuintuple(h, c) {
			return true
		}
	}
	return false
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
