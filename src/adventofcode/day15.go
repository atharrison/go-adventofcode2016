package adventofcode

import "fmt"

func Day15() {

	/*
		Disc #1 has 13 positions; at time=0, it is at position 1.
		Disc #2 has 19 positions; at time=0, it is at position 10.
		Disc #3 has 3 positions; at time=0, it is at position 2.
		Disc #4 has 7 positions; at time=0, it is at position 1.
		Disc #5 has 5 positions; at time=0, it is at position 3.
		Disc #6 has 17 positions; at time=0, it is at position 5.
	*/

	discs := []*Disc{
		//Disc #1 has 13 positions; at time=0, it is at position 1.
		&Disc{
			Number:    1,
			Positions: 13,
			Current:   1,
		},
		///Disc #2 has 19 positions; at time=0, it is at position 10.
		&Disc{
			Number:    2,
			Positions: 19,
			Current:   10,
		},
		//Disc #3 has 3 positions; at time=0, it is at position 2.
		&Disc{
			Number:    3,
			Positions: 3,
			Current:   2,
		},
		//Disc #4 has 7 positions; at time=0, it is at position 1.
		&Disc{
			Number:    4,
			Positions: 7,
			Current:   1,
		},
		//Disc #5 has 5 positions; at time=0, it is at position 3.
		&Disc{
			Number:    5,
			Positions: 5,
			Current:   3,
		},
		//Disc #6 has 17 positions; at time=0, it is at position 5.
		&Disc{
			Number:    6,
			Positions: 17,
			Current:   5,
		},

		// Part 2
		&Disc{
			Number:    7,
			Positions: 11,
			Current:   0,
		},
	}

	solved := false
	iterations := 0
	for !solved {
		solved = iterateDiscs(discs)
		iterations += 1
	}

	fmt.Printf("Iterations: %v", iterations)
}

func (d *Disc) increment() {
	d.Current += 1
	d.Current = d.Current % d.Positions
}

func iterateDiscs(discs []*Disc) bool {
	allZero := true
	for _, d := range discs {
		d.increment()
		if (d.Current+d.Number)%d.Positions != 0 {
			allZero = false
		}
	}
	return allZero

}

type Disc struct {
	Number    int
	Positions int
	Current   int
}
