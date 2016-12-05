package adventofcode

import (
	"crypto/md5"
	"fmt"
	//"os"
)

func Day05() {

	input := "uqwqemis"
	//input := "abc"

	index := 0
	var code string
	for len(code) < 8 {
		toHash := fmt.Sprintf("%v%v", input, index)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
		if string(hash[0:5]) == "00000" {
			fmt.Printf("Found %v in hash %v\n", index, hash)
			code = code + string(hash[5])
		}
		//if index % 1000 == 0 {
		//	fmt.Printf("index: %v, Hash: %v\n", index, hash[0:5])
		//}
		index++
	}
	fmt.Println(code)
}
