package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	cubeGamepart1()
	cubeGamePart2()
}
func cubeGamePart2() {
	powerArray := []int{}

	for _, v := range PuzzleInput2() {
		power := 0

		red := 0
		green := 0
		blue := 0

		firstSplits := strings.Split(v, ":")
		secondStrings := strings.Split(firstSplits[1], ";")
		// SeconsString: "3 blue, 4 red", "3 red, 4 blue" ex

		for _, x := range secondStrings {
			thirdStrings := strings.Split(x, ",")
			// thirdStrings: "3 blue", "4 red" ex
			for _, y := range thirdStrings {
				fourthStrings := strings.Split(y, " ")
				// fourthStrings: "3", "blue" ex

				number, _ := strconv.Atoi(fourthStrings[1])
				color := fourthStrings[2]

				if color == "red" {
					if red == 0 {
						red = number
					} else if number > red {
						red = number
					}
				}
				if color == "blue" {
					if blue == 0 {
						blue = number
					} else if number > blue {
						blue = number
					}
				}
				if color == "green" {
					if green == 0 {
						green = number
					} else if number > green {
						green = number
					}
				}
			}
		}
		power = (red * blue * green)
		powerArray = append(powerArray, power)
	}
	result := 0
	for _, v := range powerArray {
		result += v
	}
	fmt.Println("RESLUT:", result)
}

func cubeGamepart1() {
	ids := []int{}

	for _, v := range PuzzleInput2() {
		shouldReturn := false
		firstSplits := strings.Split(v, ":")
		secondStrings := strings.Split(firstSplits[1], ";")
		// SeconsString: "3 blue, 4 red", "3 red, 4 blue" ex

		for _, x := range secondStrings {
			thirdStrings := strings.Split(x, ",")
			// thirdStrings: "3 blue", "4 red" ex
			for _, y := range thirdStrings {
				fourthStrings := strings.Split(y, " ")
				// fourthStrings: "3", "blue" ex

				number, _ := strconv.Atoi(fourthStrings[1])
				color := fourthStrings[2]

				if checkIfMatch(number, color) {
					shouldReturn = true
					break
				}
			}
			if shouldReturn {
				break
			}
		}
		if shouldReturn {
			continue
		}
		id, _ := strconv.Atoi(firstSplits[0])
		ids = append(ids, id)
	}
	result := 0
	for _, v := range ids {
		result += v
	}
	fmt.Println("RESLUT:", result)
}

func checkIfMatch(number int, color string) bool {
	if color == "red" {
		if number > 12 {
			return true
		}
	}
	if color == "blue" {
		if number > 14 {
			return true
		}
	}
	if color == "green" {
		if number > 13 {
			return true
		}
	}
	return false
}

func PuzzleInput2() []string { // Path: 2023/day1.go
	array := []string{
		"1: 3 blue, 2 green, 6 red; 17 green, 4 red, 8 blue; 2 red, 1 green, 10 blue; 1 blue, 5 green",
		"2: 9 red, 2 green; 5 red, 1 blue, 6 green; 3 green, 13 red, 1 blue; 3 red, 6 green; 1 blue, 14 red, 6 green",
		"3: 6 red, 3 blue, 8 green; 6 blue, 12 green, 15 red; 3 blue, 18 green, 4 red",
		"4: 1 blue, 4 red; 2 blue, 6 red; 13 blue; 11 blue, 1 green, 8 red; 10 blue, 3 green, 2 red; 3 green, 7 blue",
		"5: 2 red, 1 blue, 8 green; 2 blue, 7 green, 3 red; 1 blue, 7 green, 4 red; 2 blue, 1 green, 1 red; 13 green, 1 blue",
		"6: 7 green, 1 red, 3 blue; 1 red, 4 blue; 6 green, 6 blue; 8 green, 1 red; 6 green, 1 red, 5 blue",
		"7: 10 blue, 1 green; 5 red, 8 blue, 3 green; 11 blue, 5 red, 8 green; 2 blue, 8 red, 5 green; 7 blue, 9 green; 6 blue, 2 green, 7 red",
		"8: 15 green, 8 blue, 3 red; 6 blue, 7 green, 5 red; 2 green, 1 red, 5 blue; 9 blue, 9 green, 5 red",
		"9: 16 red; 5 blue, 6 red, 9 green; 7 blue, 6 green, 2 red; 15 red, 5 blue, 3 green; 1 red, 6 green, 6 blue; 3 blue, 7 red, 5 green",
		"10: 17 green, 5 blue, 6 red; 18 green, 9 red; 4 red, 4 blue, 4 green; 10 red, 6 green, 5 blue; 8 red, 4 blue, 12 green",
		"11: 4 blue, 2 green, 5 red; 1 blue, 1 red; 9 blue, 1 green, 2 red; 4 red, 10 blue; 3 green, 4 blue, 3 red",
		"12: 4 green, 2 blue, 7 red; 4 blue, 2 green, 1 red; 7 green, 5 blue, 9 red",
		"13: 1 green, 3 red, 3 blue; 1 blue, 10 green; 2 green, 3 blue",
		"14: 7 red, 3 green, 12 blue; 5 red, 4 green, 6 blue; 13 blue, 1 red; 4 blue, 6 red, 2 green; 4 red, 3 blue; 9 red, 13 blue",
		"15: 4 blue, 5 red, 2 green; 7 red, 2 blue, 1 green; 17 red, 3 blue; 2 blue; 4 blue, 8 red",
		"16: 5 blue; 9 red, 14 green, 5 blue; 5 blue, 9 green",
		"17: 3 blue, 5 red; 6 blue, 1 green, 4 red; 7 green, 6 blue, 7 red; 1 red, 6 blue, 4 green; 6 green; 1 blue, 6 green",
		"18: 9 blue, 4 green, 2 red; 1 green, 9 red, 10 blue; 14 red, 10 green, 17 blue; 12 red, 1 green, 15 blue; 3 blue, 8 red, 2 green; 3 green, 11 red, 13 blue",
		"19: 1 blue, 3 red, 9 green; 14 green, 3 red, 2 blue; 1 blue, 8 red, 11 green; 4 blue, 3 red; 14 red, 4 green; 5 red, 8 green",
		"20: 2 blue, 3 red, 2 green; 3 blue, 2 green; 1 red, 4 green, 5 blue; 9 blue, 9 green, 3 red; 3 green, 1 blue, 4 red; 1 red, 9 green, 2 blue",
		"21: 11 blue, 6 red; 8 red; 7 red, 6 green, 11 blue; 7 green, 7 red, 11 blue; 6 red, 12 blue",
		"22: 7 green, 8 blue, 5 red; 12 green, 4 red, 2 blue; 12 green, 7 red, 11 blue",
		"23: 5 green, 2 blue, 7 red; 6 blue, 8 green, 3 red; 10 red, 5 blue; 6 green, 3 blue; 1 green, 8 red",
		"24: 17 blue, 1 green, 2 red; 2 red, 11 green, 9 blue; 6 red, 8 blue",
		"25: 3 red, 1 blue, 19 green; 1 blue, 1 green, 6 red; 6 green, 5 blue; 4 green, 2 red, 19 blue; 6 red, 19 blue, 18 green; 1 red, 4 blue, 1 green",
		"26: 3 red, 4 blue, 2 green; 2 red, 1 green, 3 blue; 14 blue, 1 green, 3 red; 5 green, 2 red, 10 blue; 9 blue, 2 red, 7 green; 15 blue, 4 green, 3 red",
		"27: 3 blue, 5 red, 2 green; 6 red, 7 blue, 9 green; 14 green, 11 red, 6 blue; 3 blue, 20 green, 3 red; 6 red, 15 green, 7 blue; 13 red, 1 blue, 14 green",
		"28: 4 blue, 7 green, 4 red; 2 red, 4 blue, 7 green; 6 blue, 11 green, 4 red; 6 blue, 6 green, 3 red; 6 green, 12 red",
		"29: 3 red, 9 blue; 5 red, 5 blue; 2 green, 3 red, 3 blue",
		"30: 1 green, 1 red, 3 blue; 1 blue, 1 red, 3 green; 1 blue; 1 blue, 3 green, 1 red; 2 blue, 2 green",
		"31: 2 blue, 1 red; 1 blue, 1 green, 1 red; 4 blue, 3 green; 1 red, 3 green, 2 blue; 2 green",
		"32: 1 blue, 6 green; 9 red, 6 green; 1 blue, 15 red, 3 green",
		"33: 18 green, 1 blue, 10 red; 10 red, 1 blue, 7 green; 11 green; 6 red, 13 green",
		"34: 10 red, 14 blue, 6 green; 2 green, 13 blue, 1 red; 8 green, 7 blue, 1 red; 9 blue, 7 green, 4 red",
		"35: 5 blue, 9 green, 2 red; 7 green, 9 blue, 5 red; 1 green, 5 red",
		"36: 10 red; 5 red, 1 green, 1 blue; 2 green, 8 red; 9 red, 2 green; 1 blue, 10 red; 6 red, 1 green, 1 blue",
		"37: 13 red, 1 blue, 7 green; 1 green, 9 red, 3 blue; 4 red, 1 blue, 11 green; 1 red; 1 red, 1 blue; 6 red, 3 blue, 2 green",
		"38: 3 blue, 12 red, 7 green; 1 green; 12 red, 1 blue",
		"39: 7 green, 12 blue, 2 red; 3 red, 10 blue, 7 green; 2 red, 8 green, 3 blue; 3 red, 12 blue, 5 green",
		"40: 7 green, 5 red; 1 green, 2 blue; 2 red, 1 green, 7 blue",
		"41: 1 red, 7 green, 2 blue; 2 green, 2 blue; 4 blue, 7 green, 1 red; 1 blue, 1 red, 7 green; 6 blue, 2 red, 3 green",
		"42: 6 blue; 4 green, 18 blue, 1 red; 10 green, 14 blue, 2 red; 6 blue, 4 green; 2 red, 13 blue, 6 green; 6 green, 1 red, 5 blue",
		"43: 5 blue, 12 red; 5 blue, 2 green, 7 red; 9 red, 4 blue; 1 green, 11 red, 2 blue; 5 red, 1 green; 2 blue, 3 red, 1 green",
		"44: 4 blue, 9 red, 4 green; 4 blue, 10 red; 4 green, 5 red; 1 green, 2 red, 3 blue",
		"45: 7 green, 2 blue, 18 red; 19 red, 7 green; 8 green, 1 blue, 19 red; 2 green, 12 red; 6 red, 5 green; 7 green, 10 red",
		"46: 1 blue, 15 red, 11 green; 7 red, 1 green, 5 blue; 13 red, 2 blue, 2 green; 7 green, 5 blue, 10 red; 12 green, 3 red, 1 blue",
		"47: 2 blue, 2 red, 5 green; 7 green, 2 red, 7 blue; 10 blue, 2 red, 8 green",
		"48: 8 green, 10 red; 6 green, 5 red; 12 green, 2 blue; 17 green, 5 red, 1 blue; 14 green, 3 blue, 16 red; 1 blue, 5 red",
		"49: 5 blue, 6 red, 12 green; 8 blue, 15 green; 4 blue, 3 green, 3 red; 6 red, 11 green, 10 blue; 9 green, 2 red, 10 blue",
		"50: 10 red, 11 green, 14 blue; 6 green, 8 blue, 17 red; 2 blue, 4 red; 6 blue, 8 green, 17 red; 17 red, 9 blue, 2 green; 13 blue, 16 red, 12 green",
		"51: 12 red, 2 green, 7 blue; 5 blue, 10 red; 1 blue, 7 red, 1 green; 14 blue, 2 red, 1 green",
		"52: 5 blue, 5 red, 8 green; 1 blue, 9 green, 7 red; 4 blue, 5 red, 3 green; 7 green, 2 blue, 2 red; 5 red, 3 blue, 17 green; 19 green, 1 red",
		"53: 4 red, 1 blue, 2 green; 1 green; 2 red; 1 blue, 2 green; 2 green, 4 red",
		"54: 7 red; 9 red, 2 blue, 14 green; 1 blue, 5 green; 7 green, 3 blue",
		"55: 11 blue, 2 green; 11 blue, 9 green, 12 red; 8 green, 6 blue, 12 red",
		"56: 2 green, 1 red, 2 blue; 4 red, 5 green, 5 blue; 5 red, 5 blue, 10 green; 8 red, 3 green, 3 blue",
		"57: 7 red, 3 green; 1 blue, 6 red, 1 green; 1 blue; 7 red, 1 green, 1 blue; 2 red, 1 blue, 1 green; 3 green, 1 blue",
		"58: 9 blue, 2 red; 2 green, 9 blue, 2 red; 5 blue, 4 green",
		"59: 8 red; 5 green, 1 blue, 3 red; 1 green, 8 red",
		"60: 9 green, 8 blue, 3 red; 10 green, 4 red; 8 blue, 2 green, 4 red; 1 red, 5 green, 7 blue; 1 green, 4 blue, 1 red; 4 blue",
		"61: 5 blue, 9 red, 4 green; 5 green, 7 blue, 6 red; 7 green, 8 red; 7 blue, 4 red, 2 green; 8 red, 4 blue, 5 green; 3 green, 9 red, 7 blue",
		"62: 9 red, 10 blue; 1 green, 7 red, 13 blue; 1 green, 11 blue; 6 red, 16 blue, 5 green; 20 red, 1 green, 3 blue",
		"63: 9 red, 8 green, 1 blue; 13 green, 12 red, 1 blue; 7 green, 5 red, 3 blue",
		"64: 3 red, 2 blue, 10 green; 3 green; 1 blue, 8 green, 2 red; 7 red, 1 blue, 4 green; 9 red, 1 blue, 4 green",
		"65: 7 red, 6 green; 1 blue, 4 green, 7 red; 6 red; 6 red, 4 green, 1 blue",
		"66: 4 blue, 4 green; 1 red, 7 green, 1 blue; 7 green, 3 red, 3 blue; 1 blue, 1 red, 6 green; 3 red, 7 green",
		"67: 5 green, 16 blue, 5 red; 4 red, 7 green, 3 blue; 4 red, 4 green, 9 blue; 12 green, 5 red; 15 green, 3 red; 10 blue, 6 red, 1 green",
		"68: 3 green, 3 blue, 5 red; 2 green, 6 blue; 2 green, 3 blue, 1 red; 1 blue, 11 red",
		"69: 5 green, 1 blue; 16 green, 9 red; 10 red, 18 green",
		"70: 1 blue, 1 green; 1 red; 1 red, 2 blue, 1 green; 1 green, 2 red; 2 blue, 2 red; 1 red",
		"71: 11 red; 2 green, 3 blue, 13 red; 1 green, 3 blue; 15 red, 1 green, 3 blue; 4 red",
		"72: 2 blue, 6 red, 18 green; 6 red, 8 green, 7 blue; 5 blue, 3 red, 12 green; 3 red, 2 blue, 4 green",
		"73: 12 blue, 7 green, 4 red; 5 green, 2 red, 4 blue; 3 green, 3 red, 10 blue; 1 green, 12 blue, 6 red; 3 blue, 6 green, 14 red",
		"74: 3 red; 1 blue, 8 green, 11 red; 3 green, 2 red",
		"75: 5 green, 2 red, 1 blue; 8 green, 2 red; 11 green, 2 red; 2 red, 17 green; 3 blue, 3 green, 2 red",
		"76: 1 blue, 5 green, 4 red; 8 green, 11 blue, 5 red; 8 blue, 2 red, 11 green",
		"77: 1 red, 11 blue, 7 green; 8 green, 4 blue; 1 blue, 8 green",
		"78: 1 green, 1 red, 1 blue; 3 green, 1 blue, 3 red; 10 green, 1 blue; 12 green",
		"79: 1 red, 11 blue, 6 green; 3 green, 3 red, 5 blue; 16 blue, 1 red, 5 green; 11 blue, 3 green, 2 red; 8 blue, 6 green, 4 red",
		"80: 5 green; 6 green, 7 red, 4 blue; 7 green, 5 blue; 6 blue, 6 green; 7 blue, 7 green; 6 green, 7 blue, 5 red",
		"81: 1 green, 14 blue; 11 blue, 1 red; 1 red, 16 green, 2 blue; 9 green, 1 red, 13 blue; 10 green, 8 blue",
		"82: 7 green, 7 red, 3 blue; 4 blue, 1 green, 4 red; 2 green, 14 blue, 3 red",
		"83: 15 blue; 2 blue, 1 green, 4 red; 8 green, 4 red, 6 blue",
		"84: 12 blue, 17 green; 6 green, 1 red, 16 blue; 1 blue, 1 red; 5 blue, 11 green",
		"85: 5 blue, 15 green, 3 red; 4 blue, 1 green, 11 red; 8 red, 2 blue, 4 green",
		"86: 11 blue, 16 green, 16 red; 11 blue, 17 red, 10 green; 8 green, 7 red",
		"87: 2 red, 4 green, 2 blue; 2 blue, 6 green; 2 red, 3 blue, 3 green; 1 red, 4 green; 1 green, 2 blue, 2 red; 4 blue, 4 green",
		"88: 10 red, 7 green; 2 blue, 6 red, 1 green; 8 blue, 8 red, 7 green; 2 green, 5 blue, 2 red; 3 blue, 3 red, 6 green",
		"89: 9 blue, 16 green; 2 red, 5 blue, 6 green; 12 blue, 15 green; 8 green, 2 red, 3 blue",
		"90: 18 red, 1 blue; 3 red, 5 blue, 4 green; 1 blue, 2 green, 6 red; 2 green, 16 red, 3 blue; 5 blue, 13 red, 5 green",
		"91: 4 red, 7 green, 1 blue; 3 green, 16 blue, 2 red; 4 green, 8 blue",
		"92: 4 red, 3 green; 5 red, 11 green, 1 blue; 16 green, 13 red; 15 green, 14 red, 3 blue; 3 red, 5 green, 2 blue",
		"93: 2 blue, 1 red, 3 green; 10 blue, 1 red, 10 green; 11 blue, 16 green, 4 red; 2 green, 20 blue, 7 red; 11 green, 8 red, 15 blue; 9 green, 10 blue, 1 red",
		"94: 2 blue, 12 red, 10 green; 16 red, 9 blue, 6 green; 5 green, 9 blue, 11 red; 4 red, 2 blue",
		"95: 2 green, 9 red, 1 blue; 2 blue, 1 red; 2 green, 5 blue, 3 red",
		"96: 1 green, 5 red, 13 blue; 1 green, 2 red, 13 blue; 2 green, 2 red, 17 blue; 3 red, 1 green; 6 red, 2 green; 1 green, 7 blue, 4 red",
		"97: 1 green, 1 red, 1 blue; 2 blue, 11 green; 1 blue, 13 green; 9 blue, 6 green, 1 red; 10 green, 8 blue",
		"98: 12 green, 9 red; 12 green, 10 blue, 3 red; 3 red, 13 green, 7 blue",
		"99: 8 green, 10 blue, 1 red; 10 green, 2 red, 6 blue; 3 green, 1 blue, 1 red; 10 blue, 1 red",
		"100: 8 blue, 6 green, 8 red; 7 red, 2 blue; 2 red, 10 green, 10 blue; 9 green, 7 red; 3 red, 7 green, 1 blue",
	}
	return array

}

func PuzzleInputEx2() []string {
	array := []string{
		"1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	return array
}
