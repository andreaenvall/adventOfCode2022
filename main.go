package main

import (
	"fmt"

	eight "github.com/andreaenvall/adventOfCode2022/dayEight"
	five "github.com/andreaenvall/adventOfCode2022/dayFive"
	four "github.com/andreaenvall/adventOfCode2022/dayFour"
	nine "github.com/andreaenvall/adventOfCode2022/dayNine"
	one "github.com/andreaenvall/adventOfCode2022/dayOne"
	seven "github.com/andreaenvall/adventOfCode2022/daySeven"
	six "github.com/andreaenvall/adventOfCode2022/daySix"
	ten "github.com/andreaenvall/adventOfCode2022/dayTen"
	three "github.com/andreaenvall/adventOfCode2022/dayThree"
	dayTwentyOne "github.com/andreaenvall/adventOfCode2022/dayTwentyOne"
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

	// Day 3
	puzzle3 := three.PuzzleInput()
	fmt.Println("Result day 3: ", three.DayThreePartOne(puzzle3))
	fmt.Println("Result day 3: ", three.DayThreePartTwo(puzzle3))

	puzzle4 := four.PuzzleInput()
	fmt.Println("Result day 4: ", four.DayFourPartOne(puzzle4))
	fmt.Println("Result day 4: ", four.DayFourPartTwo(puzzle4))

	puzzle5 := five.PuzzleInput()
	fmt.Println("Result day 5: ", five.DayFivePartOne(puzzle5))
	fmt.Println("Result day 5: ", five.DayFivePartTwo(puzzle5))

	puzzle6 := six.PuzzleInput()
	fmt.Println("Result day 6: ", six.DaySixPartOne(puzzle6))
	fmt.Println("Result day 6: ", six.DaySixPartTwo(puzzle6))

	//This one is kaos
	puzzle7 := seven.PuzzleInput()
	fmt.Println("Result day 7: ", seven.DaySevenPartOne(puzzle7))
	// fmt.Println("Result day 7: ", seven.DaySevenPartTwo(puzzle7))
	puzzle8 := eight.PuzzleInput()
	fmt.Println("Result day 8: ", eight.DayEightPartOne(puzzle8))
	fmt.Println("Result day 8: ", eight.DayEightPartTwo(puzzle8))

	puzzle9 := nine.PuzzleInputEx()
	fmt.Println("Result day 9: ", nine.DayNinePartOne(puzzle9, 4, 0, 4, 0))
	//fmt.Println("Result day 9: ", nine.DayNinePartTwo(puzzle9))
	puzzle10 := ten.PuzzleInput()
	fmt.Println("Result day 10: ", ten.DayTenPartOne(puzzle10))
	//fmt.Println("Result day 10: ", ten.DayTenPartTwo(puzzle10))
	fmt.Println("Result day 21: ")
	dayTwentyOne.MainDayTwentyOne()

}
