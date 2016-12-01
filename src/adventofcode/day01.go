package adventofcode

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	x int
	y int

	direction rune
	visited   map[string]bool // Part 2
}

func (m *Map) String() string {
	dist := math.Abs(float64(m.x)) + math.Abs(float64(m.y))
	return fmt.Sprintf("Facing %v, at [%v, %v], distance: %v\n", string(m.direction), m.x, m.y, dist)
}

func NewMap() *Map {
	return &Map{
		x:         0,
		y:         0,
		direction: 'N',
		visited:   make(map[string]bool),
	}

}

func (m *Map) Visit(x int, y int) {
	point := fmt.Sprintf("%v:%v", x, y)
	if m.visited[point] {
		dist := math.Abs(float64(x)) + math.Abs(float64(y))
		fmt.Printf("First Point visited twice is [%v, %v], distance: %v\n", x, y, dist)
		os.Exit(1)
	} else {
		m.visited[point] = true
	}

}

func Day01() {
	fmt.Println("Day01 START")

	input := readFileAsString("data/day01input.txt")
	tokens := strings.Split(input, ", ")
	fmt.Printf("tokens: %v\n", tokens)

	m := NewMap()
	for _, token := range tokens {
		fmt.Printf("Processing %v\n", token)
		m.ProcessStep(token)
	}
	fmt.Printf("Map: %v\n", m)
}

func (m *Map) ProcessStep(step string) {
	if step[0] == 'L' {
		fmt.Printf("Left")
		m.RotateLeft()
	} else {
		fmt.Printf("Right")
		m.RotateRight()
	}
	dist, _ := strconv.ParseInt(strings.TrimSpace(step[1:]), 10, 64)
	fmt.Printf(" %v\n", dist)
	//m.Move(int(dist))
	m.MoveAndTag(int(dist)) // Part 2
}

func (m *Map) Move(distance int) {
	fmt.Printf("Moving %v by %v\n", string(m.direction), distance)
	switch m.direction {
	case 'N':
		m.x += distance
	case 'E':
		m.y += distance
	case 'S':
		m.x -= distance
	case 'W':
		m.y -= distance
	}
}

func (m *Map) MoveAndTag(distance int) {
	switch m.direction {
	case 'N':
		for i := 1; i < distance; i++ {
			m.Visit(m.x+i, m.y)
		}
		m.x += distance
	case 'E':
		for i := 1; i < distance; i++ {
			m.Visit(m.x, m.y+i)
		}
		m.y += distance
	case 'S':
		for i := 1; i < distance; i++ {
			m.Visit(m.x-i, m.y)
		}
		m.x -= distance
	case 'W':
		for i := 1; i < distance; i++ {
			m.Visit(m.x, m.y-i)
		}
		m.y -= distance
	}
	m.Visit(m.x, m.y)
}

func (m *Map) RotateRight() {
	switch m.direction {
	case 'N':
		m.direction = 'E'
	case 'E':
		m.direction = 'S'
	case 'S':
		m.direction = 'W'
	case 'W':
		m.direction = 'N'
	}
}

func (m *Map) RotateLeft() {
	switch m.direction {
	case 'N':
		m.direction = 'W'
	case 'E':
		m.direction = 'N'
	case 'S':
		m.direction = 'E'
	case 'W':
		m.direction = 'S'
	}
}
