package main

import (
	"fmt"

	one "github.com/andreaenvall/adventOfCode2022/dayOne"
	two "github.com/andreaenvall/adventOfCode2022/dayTwo"
)

func main() {
	// Day 1
	puzzle := one.PuzzleInput()
	fmt.Println("Result day 1: ", one.DayOnepartOne(puzzle))
	fmt.Println("Result day 1: ", one.DayOnepartTwo(puzzle))

	// Day 2
	puzzle2 := two.PuzzleInput()
	fmt.Println("Result day 2: ", two.DayTwoPartOne(puzzle2))
	fmt.Println("Result day 2: ", two.DayTwoPartTwo(puzzle2))
}
