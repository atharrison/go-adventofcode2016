package adventofcode

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var botInstructions map[string]*BotInstruction
var valueInstructions map[int]*ValueInstruction

var bots map[string]*Bot
var outputs map[string][]int

type Bot struct {
	Number string
	Chips  []int
}

var part1 = false
var part2 = true

func Day10() {
	day := "10"
	//filename := fmt.Sprintf("data/day%vinput_sample.txt", day)
	filename := fmt.Sprintf("data/day%vinput.txt", day)
	input := readFileAsLines(filename)

	outputs = make(map[string][]int)

	botInstructions = make(map[string]*BotInstruction)
	valueInstructions = make(map[int]*ValueInstruction)
	for _, line := range input {
		fmt.Printf("Line: %v\n", line)

		if strings.HasPrefix(line, "value") {
			processValueInstruction(line)
		} else {
			processBotInstruction(line)
		}
	}

	//Initialize State
	bots = make(map[string]*Bot)

	for _, instr := range valueInstructions {
		NewBotFromValInstr(instr)
	}

	fmt.Printf("Values: %v\n", valueInstructions)
	fmt.Printf("Bots: %v\n", botInstructions)

	readyBots := findReadyBots()

	processBots(readyBots)
}

func findReadyBots() []*Bot {
	rBots := []*Bot{}
	for _, bot := range bots {
		if len(bot.Chips) == 2 {
			rBots = append(rBots, bot)
		}
	}
	return rBots
}

func NewBotFromValInstr(instr *ValueInstruction) {

	if b, ok := bots[instr.BotNum]; ok {
		b.GiveChip(instr.Chip)
		//return b
	} else {
		b := &Bot{
			Number: instr.BotNum,
			Chips:  []int{instr.Chip},
		}
		bots[b.Number] = b
		//return b
	}
}

func processBots(readyBots []*Bot) {
	for {
		bot := readyBots[0]
		readyBots = readyBots[1:]

		bot1, bot2 := bot.HandleChips()
		if bot1 != nil && len(bot1.Chips) == 2 {
			readyBots = append(readyBots, bot1)
		}
		if bot2 != nil && len(bot2.Chips) == 2 {
			readyBots = append(readyBots, bot2)
		}
	}
}

func (b *Bot) HandleChips() (*Bot, *Bot) {

	var rBot1, rBot2 *Bot

	instr := botInstructions[b.Number]

	lowChip := int(math.Min(float64(b.Chips[0]), float64(b.Chips[1])))
	highChip := int(math.Max(float64(b.Chips[0]), float64(b.Chips[1])))

	if lowChip == 17 && highChip == 61 && part1 {
		fmt.Printf("Bot: %v Had 17 and 61!!!!!\n", b.Number)
		os.Exit(0) //Bail!
	} else {
		fmt.Printf("Bot %v Handling %v and %v\n", b.Number, lowChip, highChip)
	}

	// Handle Low Chip
	if instr.LowBot != "" {
		if bot1, ok := bots[instr.LowBot]; ok {
			bot1.GiveChip(lowChip)
			rBot1 = bot1
		} else {
			bot1 := NewBot(instr.LowBot, lowChip, 0)
			bots[bot1.Number] = bot1
			rBot1 = bot1
		}

	} else {
		if output, ok := outputs[instr.LowOutput]; ok {
			output = append(output, lowChip)
		} else {
			outputs[instr.LowOutput] = []int{lowChip}
		}
		rBot1 = nil
	}

	// Handle High Chip
	if instr.HighBot != "" {
		if bot2, ok := bots[instr.HighBot]; ok {
			bot2.GiveChip(highChip)
			rBot2 = bot2
		} else {
			bot2 := NewBot(instr.HighBot, highChip, 0)
			bots[bot2.Number] = bot2
			rBot2 = bot2
		}

	} else {
		if output, ok := outputs[instr.HighOutput]; ok {
			output = append(output, highChip)
			rBot2 = nil
		} else {
			outputs[instr.HighOutput] = []int{highChip}
			rBot2 = nil
		}
	}

	fmt.Printf("rBot1: %v, rBot2: %v\n", rBot1, rBot2)
	if out0, ok0 := outputs["0"]; ok0 && part2 {
		if out1, ok1 := outputs["1"]; ok1 {
			if out2, ok2 := outputs["2"]; ok2 {
				product := out0[0] * out1[0] * out2[0]
				fmt.Printf("Output result: %v*%v*%v=%v\n", out0[0], out1[0], out2[0], product)
				os.Exit(0)
			}
		}
	}
	return rBot1, rBot2
}

func (b *Bot) GiveChip(chip int) {
	b.Chips = append(b.Chips, chip)
}

func (b *Bot) String() string {
	return fmt.Sprintf("%v=%v", b.Number, b.Chips)
}

func NewBot(num string, chip1 int, chip2 int) *Bot {
	if chip2 > 0 {
		return &Bot{
			Number: num,
			Chips:  []int{chip1, chip2},
		}
	} else if chip1 > 0 {
		return &Bot{
			Number: num,
			Chips:  []int{chip1},
		}
	} else {
		return &Bot{
			Number: num,
			Chips:  []int{},
		}
	}
}

func processValueInstruction(line string) {
	v := NewValueInstruction(line)
	valueInstructions[v.Chip] = v
}

func processBotInstruction(line string) {
	b := NewBotInstruction(line)
	botInstructions[b.Number] = b
}

type ValueInstruction struct {
	Chip   int
	BotNum string
}

func NewValueInstruction(line string) *ValueInstruction {
	tokens := strings.Split(line, " ")
	chip, _ := strconv.Atoi(tokens[1])
	return &ValueInstruction{
		Chip:   chip,
		BotNum: tokens[5],
	}
}

type BotInstruction struct {
	Number     string
	LowBot     string
	HighBot    string
	LowOutput  string
	HighOutput string
}

func NewBotInstruction(line string) *BotInstruction {
	tokens := strings.Split(line, " ")
	if tokens[5] == "output" { //low output
		if tokens[10] == "output" { //high output
			return &BotInstruction{
				Number:     tokens[1],
				LowOutput:  tokens[6],
				HighOutput: tokens[11],
			}
		} else { //high bot
			return &BotInstruction{
				Number:    tokens[1],
				LowOutput: tokens[6],
				HighBot:   tokens[11],
			}

		}

	} else { //low bot
		if tokens[10] == "output" { //high output
			return &BotInstruction{
				Number:     tokens[1],
				LowBot:     tokens[6],
				HighOutput: tokens[11],
			}
		} else { // high bot
			return &BotInstruction{
				Number:  tokens[1],
				LowBot:  tokens[6],
				HighBot: tokens[11],
			}

		}
	}
}
