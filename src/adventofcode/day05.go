package adventofcode

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func Day05() {

	input := "uqwqemis"
	//input := "abc"

	////Part 1
	//index := 0
	//var code string
	//for len(code) < 8 {
	//	toHash := fmt.Sprintf("%v%v", input, index)
	//	hash := fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
	//	if string(hash[0:5]) == "00000" {
	//		fmt.Printf("Found %v in hash %v\n", index, hash)
	//		code = code + string(hash[5])
	//	}
	//	//if index % 1000 == 0 {
	//	//	fmt.Printf("index: %v, Hash: %v\n", index, hash[0:5])
	//	//}
	//	index++
	//}
	//fmt.Println(code)

	//Part 2
	//index := 0
	//code := make(map[int]string)
	//for len(code) < 8 {
	//	toHash := fmt.Sprintf("%v%v", input, index)
	//	hash := fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
	//	if string(hash[0:5]) == "00000" {
	//		if hash[5] >= byte('a') {
	//			index++
	//			continue
	//		}
	//		pos64, _ := strconv.ParseInt(string(hash[5]), 10, 64)
	//		pos := int(pos64)
	//		if _, ok := code[pos]; !ok {
	//			if pos < 8 {
	//				code[pos] = string(hash[6])
	//				fmt.Printf("Found %v in hash %v\n", index, hash)
	//				fmt.Printf("Putting %v in space %v\n", string(hash[6]), pos)
	//			}
	//		}
	//	}
	//	index++
	//}
	//fmt.Println(code)
	//for i := 0; i < 8; i++ {
	//	fmt.Printf("%v", code[i])
	//}

	//Part 2, With animation
	index := 0
	code := make(map[int]string)
	for len(code) < 8 {
		toHash := fmt.Sprintf("%v%v", input, index)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
		if string(hash[0:5]) == "00000" {
			if hash[5] >= byte('a') {
				index++
				continue
			}
			pos64, _ := strconv.ParseInt(string(hash[5]), 10, 64)
			pos := int(pos64)
			if _, ok := code[pos]; !ok {
				if pos < 8 {
					code[pos] = string(hash[6])
				}
			}
		}
		index++
		if index%1000 == 0 {
			fmt.Printf("\r")
			for i := 0; i < 8; i++ {
				if _, ok := code[i]; ok {
					fmt.Printf("%v", code[i])
				} else {
					fmt.Printf("%s", string(hash[i+9]))
				}
			}
		}
	}
	fmt.Printf("\r")
	for i := 0; i < 8; i++ {
		fmt.Printf("%v", code[i])
	}
	fmt.Println()

}
