package main

import (
	"fmt"

	one "github.com/andreaenvall/adventOfCode2022/dayOne"
)

func main() {
	// Day 1
	puzzle := one.PuzzleInput()
	fmt.Println("Result: ", one.DayOnepartOne(puzzle))
	fmt.Println("Result: ", one.DayOnepartTwo(puzzle))
}
