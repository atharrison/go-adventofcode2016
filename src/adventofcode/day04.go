package adventofcode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Room04 struct {
	Letters            []string
	SectorId           int
	Checksum           string
	CalculatedChecksum string
	Counts             map[string]int
	RankedPairs        PairList
	DecryptedWords     []string // Part 2
}

func NewRoom04(line string) *Room04 {
	tokens := strings.Split(line, "-")

	length := len(tokens)
	letters := tokens[:length-1]

	last := tokens[length-1]
	lastSplit := strings.Split(last, "[")
	sectorId, _ := strconv.ParseInt(lastSplit[0], 10, 64)

	r := &Room04{
		Letters:  letters,
		SectorId: int(sectorId),
		Checksum: lastSplit[1][0 : len(lastSplit[1])-1],
	}
	return r
}

func (r *Room04) Calculate() {
	counts := make(map[string]int)
	for _, item := range r.Letters {
		for i := 0; i < len(item); i++ {
			counts[string(item[i])]++
		}

	}
	r.Counts = counts
}

func (r *Room04) RankByLetterCount() {
	pl := make(PairList, len(r.Counts))
	i := 0
	for k, v := range r.Counts {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	r.RankedPairs = pl
}

func (r *Room04) CalculateChecksum() {
	var letters string
	for i := 0; i < 5; i++ {
		letters += r.RankedPairs[i].Key
	}
	r.CalculatedChecksum = letters
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		//fmt.Printf("comparing letters. %v less than %v? %v\n", int(p[i].Key[0]), int(p[j].Key[0]), int(p[i].Key[0]) < int(p[j].Key[0]))
		return int(p[i].Key[0]) > int(p[j].Key[0])
	}
	return p[i].Value < p[j].Value
}
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

//Part 2
func (r *Room04) Decrypt() {
	r.DecryptedWords = make([]string, len(r.Letters))
	for w, word := range r.Letters {
		var newWord string
		for i := 0; i < len(word); i++ {
			newLetter := (int(word[i]) + (r.SectorId % 26))
			if newLetter > int('z') {
				newLetter = newLetter - 26
			}
			newWord += string(rune(newLetter))
		}
		r.DecryptedWords[w] = newWord
	}
}

func Day04() {
	day := "04"
	filename := fmt.Sprintf("data/day%vinput.txt", day)
	input := readFileAsLines(filename)

	rooms := []*Room04{}
	total := 0
	for _, line := range input {
		//fmt.Printf("Line: %v\n", line)
		room := NewRoom04(line)
		room.Calculate()
		room.RankByLetterCount()
		room.CalculateChecksum()
		//fmt.Println(room.Checksum)
		if room.Checksum == room.CalculatedChecksum {
			total += room.SectorId

			// Part 2
			room.Decrypt()
			fmt.Printf("%v\t%v\n", room.SectorId, room.DecryptedWords)
			// Grep the output for 'northpole object storage'
		}

		rooms = append(rooms, room)
		//fmt.Println(room)
	}
	fmt.Printf("Total: %v\n", total) // Part 1
	//fmt.Printf("%v\n",input)
}
