package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	findWinningNumbers()
	findAllCopiesOfCard()
}
func findAllCopiesOfCard() {
	cards := PuzzleInput4()
	Wincount := 0

	scratchcards := 0
	for i := 0; i < len(cards); i++ {
		scratchcards++
		// isfirst := true
		numbers := strings.Split(cards[i], ":")
		winLoseNum := strings.Split(numbers[1], "|")

		winningNumb := winLoseNum[0]
		myNumb := winLoseNum[1]
		id := strings.Split(numbers[0], " ")

		idInt, _ := strconv.Atoi(id[1])

		winningNumbsArr := strings.Split(winningNumb, " ")
		myNumbsArr := strings.Split(myNumb, " ")

		for _, wN := range winningNumbsArr {
			if slices.Contains(myNumbsArr, wN) {
				if wN != "" && wN != " " {
					Wincount++
				}
			}
		}
		if Wincount != 0 {
			for j := 1; j <= Wincount; j++ {
				index := idInt - 1 + j
				cards = append(cards, cards[index])

			}
		}
		Wincount = 0
	}

	fmt.Println("Cards", len(cards))
}

func findWinningNumbers() {
	winningValue := 0
	winningValueCard := 0
	cards := PuzzleInput4()
	count := 0
	// firstKey := -1

	for _, card := range cards {
		isfirst := true
		count++
		numbers := strings.Split(card, ":")
		winLoseNum := strings.Split(numbers[1], "|")

		winningNumb := winLoseNum[0]
		myNumb := winLoseNum[1] // Card 1

		winningNumbsArr := strings.Split(winningNumb, " ")
		myNumbsArr := strings.Split(myNumb, " ")

		for _, wN := range winningNumbsArr {
			for _, mN := range myNumbsArr {
				if wN == mN && wN != "" && mN != "" {
					if isfirst {
						fmt.Println(" fisrt", winningValue)
						winningValueCard += 1
					} else {
						fmt.Println("else fisrt", winningValue)
						winningValueCard += winningValueCard
					}
					isfirst = false
				}
			}

		}
		fmt.Println("Winning value per card: ", winningValueCard)
		winningValue += winningValueCard
		winningValueCard = 0
	}
	fmt.Println("Winning value: ", winningValue)
}

func PuzzleInputEx4() []string {
	array := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}
	return array
}

func PuzzleInput4() []string {
	array := []string{
		"Card 1: 79  1  6  9 88 95 84 69 83 97 | 42 95  1  6 71 69 61 99 84 12 32 96  9 82 88 97 53 24 28 65 83 38  8 68 79",
		"Card 2: 34 76 23 61 56 74 13 42 18  6 | 18 13 21 64 74 97 34 43 31 23 56 82 76 61 45 69 10 81 48  6  9 30 47 95 42",
		"Card 3: 12 88 28 50 46 69 62 95  6 51 | 66 12 62 82  6 46 77 88 36 74 50 54 40 99 89 11 33 78 87 69 75 96  2 21 71",
		"Card 4: 50 80 66 43 82 53 35 51 39 48 | 43 82 48 10 91  7 80 66 51 63 84 35 19 44  9 39 72 85 50 53 73  1 26 75 86",
		"Card 5: 63 83 42 98 60 47 36 59 93 18 | 53 47 67  5 17 60 92 93 20 84 10 98 39 86 41 16 31 83 42 94 25 82 61 95 44",
		"Card 6: 54 26 43 52 75 55 84 71 85 69 | 99 80 79 24 64 13 93  1 87  9 15 50 32 89 72 52 60 53 12 96 30 98 86 31 92",
		"Card 7:  4 98 40 78 24 38 80 32 75 97 | 32 24 78 97 57 40 82 75 70 76 98 37 35 38 64 91 28 46 39 90 80 87 22  4 45",
		"Card 8: 61 98 49 16  4 24 32 79 23 28 | 56 27 24 40 47 98 75 49 34 23 80 61 68 32 94 16 79 31  4  8 62 46 53 28  7",
		"Card 9: 77 15 52 60 86 19  5 33 65 83 | 67 69 94 16 13 17 27 53 75 59 36 38 80 86 58 61 63 24 85 44 99 70 92 47 14",
		"Card 10: 44 58 28 36 72 78 96 69 95 56 | 69 81 29 61 95 52 18 28 72 96 23 75 44 77  2 87 11 36 56 39 73 26 98 78 58",
		"Card 11: 37  2 54 40 96 82 17 67 72 95 | 16  9 92 36 54 61 72 45 13 47 96 40  2 21 82 17 76 95 37 42 67  1 88 69 24",
		"Card 12: 95 46  8 27 41 34 82  4 84 59 | 52 39  7  3 54 57 29  1 21 89 75 33 14 94 36 15 60 40 16 80 35 83  9  5 87",
		"Card 13: 67 51 83 20 42 70 24 13  1  3 |  5 38 23 32 84 34 19 62 11 69 50 17 94 10 25 86 61  3 89 30 33 63 13  2  8",
		"Card 14: 78 62 32 68  8 66 98 23 57 82 | 40 32 52 70 98 33 53 36 19 83 82 99 68 66 78 24 11 57 81 26 74 25 88 73 18",
		"Card 15: 39 31 45 89 16 57 40 50 33 51 | 87 49 24 83 54 55 39  4 53 81 19 61 20 18 97 30 52 58 13 67 50 94 47 89 56",
		"Card 16: 34 59 92  5 33  6  1 37 40 26 | 57 89 20 36 46 91 62 94 33 56 31 75 22 83  9 82 58 52 77  7 54 41 43 71 79",
		"Card 17: 37  5  8 67 99 36 56 52 44 46 | 85 65 40 10 58 23 16 11 90 94 82 12 88 96 71 77 54 78 26 19 38 43 73 21 17",
		"Card 18: 61 64 46 14 15 36 32  7 76  3 | 42 12 27 78 56 97 76 18 51 15 82 10 57 69  2  8 99  4 59  6 40 14 52 58  7",
		"Card 19: 89 21 81 73 15 58  7 25 33 67 | 84 36 88 97 95 19 94 30 70 55 96 59 34 46 52 56 22 73 50 83 86 61 25 63 16",
		"Card 20: 64 22 18 55 73 14  1 96 92 41 | 27 11 39 42 84  2 48  4 37 23 60 52 86 36 62 94 82 20 77 73 32 40 83 69 90",
		"Card 21: 23 64 38  3 13 10 35 66 72 46 | 52 67 74 57 58 83 40 29 79 48 95 24 11 22 77 62 47 30  1 71 89 96 55 86 97",
		"Card 22: 33  7 39 61 29 95 90 23 59  4 | 50 82 94 18 28 34 85 57 91 75 13 46 55 68 84 89 21 88 87 92 80 52  6 86 71",
		"Card 23:  3 30 63 96 50 67 69 52 33  9 |  3 69 78 79 96 50 64 30 33 86 81  4 77  8  7 63 97 93 59 91 45  9 22 67 52",
		"Card 24: 62 90 27 28 20 63 65 68 50 44 | 23 88 27 59 79 99 87 49 94 51 80 16  6 73 90 13 37 57  8 12 52 76 10 75 18",
		"Card 25: 75 38 28 91 95 17 30  6 68 44 | 30 76 86  2  6 35  3 45 14 82 75 66 21  5 43 94 90 38 17 72 10 24  8 19 44",
		"Card 26: 69 66 35 37 24 15 92 41  9 31 | 47 21 54 25 50 18 64 44 31 95 67 73  4 91 74 38 53 90 52 56 97 29 75 98 46",
		"Card 27: 59 72 44 21 42 89 13 38  7 61 | 72 38  8 15 17 46 85 21 62 13 61  6 44 73 59 41  7 89 12 39 23 77 32 84 42",
		"Card 28: 10 83 96 81 71 62 76 33 38 48 | 44  7 34 66 56 42 71 38 47 76 12 51 36 92  4 26 68 67 17 21 73 60 11 70 81",
		"Card 29:  5 11 24  9 71 92 77 19  3 69 | 61  3 26 39 75 70  5 58 24 23 29 89 71 33  2  9 43 69 77 19 11 38 60 92 93",
		"Card 30: 46 62 55 48  1 89 84 20 12 87 | 27 99 84 25 63 73  8 80 47 93 65 28 64 79 56 40  6 38 62 55 32 78 75 37  3",
		"Card 31: 75 31 40 20 99 72 49 71 24 47 | 84 60 26 91 20 77 59 34 88 78 18  5 13 46 19 96 86  9 97 29 61 53 54 65  8",
		"Card 32: 50 49 23 14 34 94 58 28 45 20 |  1 93 38 82 42 41 17 54 76 71 75 15 87 35 90  5 13 18 32 21 39 22  7 99 67",
		"Card 33: 61 45 57 60 80 65 94 13 16  8 | 17 57  5 31 68 32 19 23 96 16 63 24 65 35 21 85 97 70 53 48 36 60 45 40 88",
		"Card 34: 42 56 64 13 99 93 36  5 29 95 | 65 91 41 13 67 73 77 89 64 60 20 86 32 82  4 11 37 30  5 72 42 59 27 84 55",
		"Card 35: 92 59 53 87  4 65 22 11 70 26 | 67 53 42 48 99 52 63  5 80 44 27 10 41 81 82  4 31 36 61 88 33 50 76 98 38",
		"Card 36: 99 36 53 37  8 56 31 79 72 87 | 66 50 71 41 31 60 87 86 36 28 73 92 14 12 32 82 75 11  7  3 30  4 45 18 76",
		"Card 37: 44 71 49 42 99 40 84 62 46 35 | 36 23 76 78 42 47 99 97 43 80 37 73 17  4 87 88 71 53  3 60 59 44 48 62 39",
		"Card 38: 43 39 31 13 70 50 65 79 49 83 | 63 49 82 10 33 83 70 35 85 50 11 45 18 43 77 96 79 56 75 67 80 93  5 92 89",
		"Card 39:  3 47  1 17 99 74 97 29 11 18 | 70 11 20 90 54 51 94 80 17 84 32 15 63 53  4 21 50 46 88  8 48 27 39 65 98",
		"Card 40:  6 65 18  9 20 43 28  8 73 57 | 82 21 67 68 38 62  5 52  4 89 75 30 87 74 49 12 81 98 29  8 26 60 11 83 47",
		"Card 41: 20 83 52 71 96  9 78 66 79 64 | 55 91 36 83 29 51 60 37 11 99 85 67 86 77  8 80 40 65 14  5 25 56 93 90 73",
		"Card 42: 15 58 20 36  9 99 47 59 65 85 | 18 79 54 34 44 24  2 30 43 26 87 41 15 32 19 42  7 72  8 14 94 53 71 28 12",
		"Card 43:  2 23 55 36 59 77 37  6 15 50 | 79 99 83 26  1  3 40 56 33 35 12 32 51 70 64 81 78 27 25 60 95  9 21 10 65",
		"Card 44: 76 65 53 88 36 70 80 52 41 26 | 89 77 84 64  8 97 23 67 54 19 96 12 78 61 83 58 73 86 87 21 14 60 20 34 49",
		"Card 45: 69 19 18 42 25 36 83 23 24 89 | 85 56 58 89 30 18 77 42 28 91 37 70 25 65 50 55 12  9 49 61 98 45 71 72 67",
		"Card 46: 78 14  5 20  3 30 82 37 75 34 | 30 59 65 38  4  7 27 54 58 84 16 43 92 95 61 48 53 10 97 15 96 41 12 67 46",
		"Card 47: 50 69 85 23 38 44  7 68 57 64 | 37 36 15 78 14 40 73 76  2 26 90 25  4 44 89 82 21  1 68 48 95  8 54 51 17",
		"Card 48: 27 21 52 20 83 98 53 48 86 54 | 92 52 54 73 63 87 53 79 48 64 28 21 27 40 16 20 43 86  6  8 97 98 83 29 72",
		"Card 49: 62 91  8 76 11 73 13 68 86 88 | 44 16 86 25 82 53 76  2 51 68 13  8 48 28 90 54 20 91 11  6 62 17 88 73 30",
		"Card 50: 10 44 52 13 92 65 85 29 49 98 | 86 96 99  3 97 93 70 72 78 65 16 77 51 57 87 53 39 48 50 36 19 47 23 27 37",
		"Card 51: 26 30  2 44 96 77 95 51 20 32 | 52 95 83 51 93 69 30 81 74 20 79 49  4 57  2  3 92 33 29 40  1 44 86 55 22",
		"Card 52:  4 46 18 37 94 63 42 26 57 33 |  4 15 60 42 43 77 79 26 71 59 78 41 81  8 37 30 94 57 31 18 68 63 76 13 65",
		"Card 53: 54 68 97 57  3 92 13 41 40 18 | 17  4  8 86 53 93 21  1 11 52 77 38 80 59 10 45 82 27 83 67 26 49 71 31 84",
		"Card 54: 94  2  7 97 65 40 51 68 69 27 | 30 84 76 69 57 17 26 97 50  2  7 40  6 94 41 27 54 51 36 31 25 29 35 10 78",
		"Card 55: 18 48 77 41 50 93 67 28 71 47 | 60 96 49 93  2 77 58 24 79 57 44 21 16 28  6 15 12 92 66 34 10 23 31 94 69",
		"Card 56: 61 24 27  9 67 36  7 30 60 66 | 97 33 32 41 17 66 35 27 68 12 11 75 88 43 10 83 51 31  1 39 20 70 98 48 65",
		"Card 57: 40 20 83 37 22 76 23 29  7 54 | 43 31 93 58 53 67  1 66 69 99 18 57 95 81 44 62  6 45 78 92 63 14 37 50 32",
		"Card 58: 86 32 52 72 78 94 49 16 69 10 | 96 49 77 54 46 53 79 80  5 48 81 10 91 69 76 42 88 92 52 31 74 21 94 43 78",
		"Card 59: 13 68 62  3 77 30 65 87 42 12 | 68 80 33 26 77 24 76 91  5 97 43 74 16 57  1 49 62 27 17 79 78 94 45 13 42",
		"Card 60: 94 36 71 90  1 80  4 12 92 79 | 53 14 13 83  6  2 99 48 61 17 56  7 23 38 34 50 24 78 73 28 51 40 92 49 65",
		"Card 61: 27 65 72  9 13 82  2 66 81 56 | 81 19 34 70 53 85 17  6 12 72 62 32 21 29 89 63  4 43 80 77 36 90 49 56 82",
		"Card 62:  3 43 72 93 10 84 23 79 26 58 | 18 44  9 85 14 92 94 52 41 62 87 17 80 88 63  2 68 13 50 86 34 23 61 42 39",
		"Card 63: 59 84 11  4 68 27 44 73 16 64 | 43 19 81 30 35 65 56 66 22  9 54 13 21 47 82 45 79 76 34 29 36 92 91 23 17",
		"Card 64: 40 61 22 50 33 25 69  7 99 67 | 14 48 96 88 58 34 13 21 76 77 65 71  9 74 31 39 59  2 53 72 44 83  8 49 68",
		"Card 65: 79 90 74 65 40 14 37 41 16 23 | 93 85 54 36 43 73 21 27  7 69 81 38 47 99 76 26 63 87 72 55 56 97  4 11 28",
		"Card 66: 75 91 68 41 52 93 98 31 14 45 |  2 85 49 32 80 43 12 71 53  9 50 37 13 66 77  3 38 36 21 28 40 51 17 76 79",
		"Card 67: 15 95 98 23 82 38 97 81 17 37 | 95 62 37 51 33 82 56 97 45 38 16 78 17  1 10 15 23 65 47 58 11 24 81 93 98",
		"Card 68: 22 37 51 65 73 15 92 28 79 64 |  5 97 43 19 26 27 84 18 28 98  3 14 59 63 10 77 11 40 85 31 87 90 20  6 71",
		"Card 69: 25 37 33 95 31 80 69 57 21 53 | 31 69 99  9 37 90 80 72 14 50 16 74 66 21  3 95 45 53 33 65 25 83 61 57 20",
		"Card 70: 49 13 69 15 62 74 83 76 53 56 | 81 88 98 35 13  6 23 65 70 28 19 60 26 42  9  8 89 11 83  3 74 85  7 87 67",
		"Card 71: 13 98 41 95  7 99 66 58 65 86 | 41 87 11 63 97 48 13 98 99 55 65 39 70 96 79 20 58 27 61 95 66 23  7  5 86",
		"Card 72: 34  2 71 56 75 74 21 23 91 60 |  5 17 55 60 85 24 74 11 37 30 19 99 31 82 54 84 42  2 32 71 29 90  7  6 56",
		"Card 73: 18 77 29 48 78 15 92 30 17 47 | 75 99 58 70 15 78 42 89 48 63 56 92 52 85  5 18 30 74 59 77 29 17 47 41 93",
		"Card 74: 62 41 98 20 55 33  1 80 49 92 | 82 16 92 51 98 20 33 17 41 34 31  8 72  1 55 56 37 89 47 49 76 62 93  7 86",
		"Card 75: 19 55 77 40 26 74 70 61 38 89 | 69 47  6 10 89 85 96 26 74 28  2 87 55 95 24 61 19 46 38 21  5 72 67 70 76",
		"Card 76: 69 79 65 33 60 88 42 84  8 32 | 69 88 64 33 60 13 14 36 17 35 22 78 50 82 18 65 61 42 91 20 70 79 32 12 84",
		"Card 77: 49 67 77  4 42 11 15 59 65 88 | 99 50  3 74 33 11 38 85 64  6  8 93 73 68 89 79  2  5 69 14 18 97 13 56 91",
		"Card 78: 20 93 62 24 50 72 13 75 45 96 | 92 13 33 65 24 85 79 72 93 91 20 40 84 75 81  2 50 15 26 96 51 48 44 62 49",
		"Card 79: 72  2 39 78 81  1 16 68 10 84 |  6 28 12  1 72 89 65 16 20 53 82 69 27 60 68 67 45 29 61 97 76 21 30  5 47",
		"Card 80: 75 48  9 23 82 62 33 77 19 16 | 77 58 23 85  2 75 74  9 13 16 40 25 33 93 70 86  5 62 12 41 37 46  7 94 82",
		"Card 81: 44 99 69 38  7 73 25 19  8  3 | 99 98 41 55 44 58 73  2 36 89 29 71 39 77 83  8 12 86 84 30 70 69 43 85  3",
		"Card 82: 69 22 36 77 72 67 10 34 51 31 | 14 70 26 18 37 21 45 29 11 12 35 57 39 66 48 98 81  7 89 79 53 47  6 36 38",
		"Card 83: 95 57 78 99 92 43 75 89 80 32 | 77 49 63 91 64 30 69 39 24 85 28 29 40 11 10 73  1 52 46 12 83 71 44 97 98",
		"Card 84: 96 76 53 27 87  6 86 23 95 13 | 57 89 81 65 58 21 24 12 14 71 38 84 36 55 95 43 28 78 18 82 17 37 77 25  9",
		"Card 85: 69 12 27 30 47 79 31 51 37 35 | 24 23 17 89 32 47 36 59 34 83 78 39 60 72 48 22 15 66  3 38 31 70 99 19 74",
		"Card 86: 51 13 54 23 14 12 22 18 96 82 | 60 79 57 93  3 33 19 28 98 87 66 95 16  5  1 30 64 47 74 65 39 72 29 70  2",
		"Card 87: 92  5 97 85 76 61 40 42 63 67 | 98 46  7 62 99 94 70 60 29 18 38 96 75 61 78 72 79 74 49 89 43 15 83  4  6",
		"Card 88: 39 58 95 17 28 27 48 74 62 73 | 51 76 13 16 45 75 22 14 69 61 40  7 68 96 30 49 84 63 85 86 12 24 87 80 50",
		"Card 89: 27 41 39 48 99 54 36 88 46  6 | 27 45 41 74 54 46 76  8 12  6 32  5 31 39 40 36  9 20 88 48 99 52 83 50 56",
		"Card 90: 36 66 59 75 54 62 92 99 46 68 | 19 20 99 85 68 31 83 28 29 77 48 54  5 62 46 66 43 59 94 95 32 92 75 36 41",
		"Card 91: 10  3 75 22 98 53 23  2 11 84 | 11 22 46 50 32 42  4 21 74 53 87 98 10 31 81 34 38 75  3 84 70 47  2 37 23",
		"Card 92: 81 31 76 73 64 56 38 78 62 54 | 97 12  3 33 35 18 74 39 49 96 21 52  9 36 22  4 82 59 87 69  8 92 20 90 65",
		"Card 93: 19 36 14 79 72 69 75 48 65 55 | 60 58 19 64 57 48  3 14 23 75 86 20 71 46 79 55  4 69 42 12 36 72  5 65 39",
		"Card 94: 38 54 34 55 12 73 14 93 33 82 | 51  6 31 41 50 43 22 34 45 27 64 46 76 26  1 66 68 30 81 37 40 18 83 74 69",
		"Card 95: 33 36 27 35 19 40 26 31 11 61 | 77 10 23 14 61 96 91 36 33 35 40 94 60 31 29 27 11 42 26 85 19 39 53 73 80",
		"Card 96: 32 18 54 82 94 83 24 21 41 72 | 80  3 88 86 50 29 87 42 93 13 32 71 62 51 83 31 68 23 41 40 34 12 85 91 20",
		"Card 97: 58 53 51 29 22 35 86 17 71 40 | 80 60  7 22 11 40 28 85 65 26 17 21 58 86 41 52 53  3 69 51 25 63 29 71 35",
		"Card 98: 55 73 67 88 41 44  9 95 14 10 | 32 75 93 26 59 79 77 73 19 71  9 18 90 33 84 80 10 15 95 21 62 34 58 37 81",
		"Card 99: 98 27 60 28  7 45 25 19 82 76 | 90 64 23  4 32 67 45 37 18  7 65 61 78 25 14 28 81 39 48 69  8 66 60 82 76",
		"Card 100: 99 49 31 54 95 60 76 15 24 41 | 54 41 99 15 57 12 79 35 47 56 59 25  3 20 28 55 50 52  5 98 97 74 82 27 13",
		"Card 101: 15 23 33 95 74 67 93 82  9 29 | 85 47 38 94 12 22 21 59 48 58 27 18 71 52 72 68 49 10 24 23 98 16 69  5 19",
		"Card 102: 22 67 23 77 12 91 58  9 65 68 | 82 23 17 22  8  4  9 58 20 11 29 87 18 67 62 52 34 91 25 99 80 79 53 12 77",
		"Card 103: 48  3 59 50  4 61 75 58 90 80 | 19 37 59 89 61 80 48 81  4 30 55 57 25 46 11 58 90 75 68 40 35 27 50 85 86",
		"Card 104: 88 36 91 27 87 79 83 60 61 32 | 69 13 92 23 21 86 60 14 43 33 27 99  6 36 19 88 40 94 42 31 93 84 39 17 28",
		"Card 105: 42 71 90  5 38 96 21 86 37  9 | 97 59 87 68 20 79 83 37  6 96 72 40 56 48  7 67 23 47 39 18 82 49 99 42 71",
		"Card 106: 30 60 13 26 23 22  6 56 58 10 | 99 41  1 64 13 31  6 79 26  5 33 63 82 20 30 10 11  4 47 86 19 60 58 53 83",
		"Card 107: 63 40 14 21 89 38 25  9 41 82 | 39 21 40 88 48 33 87  4 36 64 95 45 44 89 50 62 38 10 17  6 78 60 91 92 71",
		"Card 108: 51  5 36 53 42 61 97 76 19 56 | 83 22 10 41 91 66  4 89 13  1  9 85 96 57 39 43 30 74 87 64 16 20 97 40 37",
		"Card 109: 82 18 33 35 16 54 71 79 22 11 | 45 26 48 22 79 66 52 28 57 99 39 13 63  5 58 55  4 97 81 29 60 77 46 47 33",
		"Card 110:  7 71 38 58 76 18 37 20 15 60 | 33 62 84 36 73 90 70 20 26 15 49 85 74 14 61 79 46 50 27 53 51 98 56 99 28",
		"Card 111: 90 87 56 91 82  5 77 49 66 21 | 86 90 28 76 29  7 48 32 81  4  8 41 79 23 74 56 96 46  9 20 16 37 26  1  2",
		"Card 112: 34 90 81 65  6 87 10 48 82 45 | 46 36 53 93 91 55 71 89 12  8 95 74 85 33 31 81 79 50 69 19 35 99 44 21 43",
		"Card 113: 47 72 45  8 63 54 93 62  7 17 | 85  6 51 90 73 94 53 80 13 83  9 91 60 43 50 23 78 52  3 61 16 99 71 68 40",
		"Card 114:  8 83 58 13 20 99 68  6 10 40 | 46  8 88 41 17 34 20 28 51 68 93  3 58 96 22 77 29 47 10 98 75 50 59 23 49",
		"Card 115: 57 44 68 95  1 52 22 30 72 54 | 88 30  1 25 52 75 68 37 54 71  4 59 50 99 48 72  9 33 44 22 95 57 13 19 63",
		"Card 116: 62 36 94 55 39 63 28 43 77 24 | 43 65 91 67 98 28 96 41 56 30 79 88 90  2 68 10 94 23 81 95 13 44 15 24 62",
		"Card 117: 80 21 34  9 19 46 59 72 88  5 | 93 88 65 18 46 49 43 75 95 19 21 25  5 67 77 86 83 34 26 80 59 36 51  9 72",
		"Card 118: 76 37 97 67 39 78 91 55  3 48 | 37  3 26 35 44 91 48 15 97 80 74 55  9 65 19 78 56 59 10 79 81 22 39 67 76",
		"Card 119: 70 76 13 72 41 11 86 84 45 59 | 35 10 22 45 41 72 70 66 25 75 14 84 57 13 30 74 95 76 47 44 86 48 11 59 79",
		"Card 120: 98 61 46 89 16 56 49 94 18 36 |  2 36 97 45 82 98 52 94 61 16  7 28 92 24 18 68 62 41 80 79 89 25 29 14 88",
		"Card 121: 23 61 52 28 74 58 55  2 77 82 | 55 32 46 97 63 45 52 38 11 23  2 77 74 75 13 44 58 51 20 84 82 69  9 28 80",
		"Card 122:  8 97 46 61 59 51 42 73 49 24 | 88  4 31 58 29 80 81 20 85 41  9 62 79 24 72 33 65 90 21 98 93 38 55  6 15",
		"Card 123:  6 19 64 72 71 78 92 27 57 93 | 71 25 46 15 47 76  6 29 27 54 96 78 11 45 97 19 17 58 92 64 93 72 41 57 63",
		"Card 124: 65 94  8 90 12  9 39 99 62 32 | 54 31 94 22 41 74 30 69 89 98 80 57 65  9 83 49 62 92 66 93 61 36 44 48 18",
		"Card 125: 43 46 16 53 13 64 11 78 47 31 | 19 54 16 43 86 34 31 46 35 62 37 53 23 72 97  1 29 18 26 69 66 12 95 15 39",
		"Card 126: 49  8 94 46 11 13 44 68 55 72 | 93  3 78 46 49 70 47  5 33 51 38 16 98 50 62 19 14 44 11 94 92 13 88 32  8",
		"Card 127: 60 94 45 91 51  8 48  6 70 34 | 50 55  8  2 38 14 83 60 34 21 71 68 79 51 56 47 91 67 16 18  6 94 48 45 70",
		"Card 128: 73 72 55 77 51 52 56 85 10 69 | 46 82 64 33 26 98 63 72 65 34 44 38 51 56 35 31 20 17 30 88 87  5 80 23 55",
		"Card 129:  4  2 93 92 24 85 96 23  5 35 | 70 50 60 18 61  2 14 17 71 36 15 20 42 98 12 19 81  4 86 40 48 28 73 54 22",
		"Card 130: 82  1 15 98 36 56 90 44 64 62 | 72 36 22 45 43 58 97 41 34 25 31 66 91 16 93 15 67 92 89 59 26 61 23  2 54",
		"Card 131:  2  3 69 28 10 81 88 45 98 35 | 75 99 97 37 94  4 29 63 16 77 95 31 19 59 46 72 81 12 40 15 48  3 53 64 24",
		"Card 132: 31 63 10 22 24  4  6 52 68 17 | 85 73 52 32 44 15 65 96 58 55  7 28 75 97 62 86 24 27 47 74 63 48  9 16 33",
		"Card 133:  3 64 77 78 39 75 57 43 69 30 | 19 70 72 14 79 96  5 56 29 88 46 40 77 57 35 90 95 47 73 48 98 99 52 50 37",
		"Card 134: 97  7  3 72 96 92 49 26 77 81 | 53 74 65 72  8 10 42 55 83 44 28 38 79 30 91 27 66  1 89 21 41 56 16 88 13",
		"Card 135: 47 90 63 12 73 13 62 74 34 30 | 22 56 49 68 19  8 60 32 65 77  4 83 69 40 46 17 26 94 42  3 29 98 97 89 80",
		"Card 136: 34 77 18 47 91 48 45 96 32 42 | 51 19 86 79 69 43 84 90 75 53 76 16  8 23 37  4  2 87 89 77 29 35 24 36 39",
		"Card 137: 46 30 37  7 81 87 74 26 96 60 | 92 40 44 91 70 55 31 47  2 97 32  4 43 34 27 50 68 64 51 80 78 52 22 38 82",
		"Card 138: 55 96 22 88 70 51 26 50 98 69 | 43 87 85 51 55 60 96 40 47 91 93 15  4 70 46 97  3 18 50 69  2 98 26 22 88",
		"Card 139: 20 71 84 27 24  7 53 34  4 72 | 48 83 31 22 42 51  8 43 29 64 57 40 46 13 68 33 16  5 50 18 12 81 55 75 70",
		"Card 140: 39 11 56 42 92 64 14 37 18 82 | 18 53 50 45 10 65 47 36 99 23 82 88 26 64 20 96 78 32 70 13 86 24 48  9  7",
		"Card 141: 45 43 61 54 19 39 70 75 74 73 | 43 75 25 78 39 74 54 10 27 61 73 34 85 71 18 38 87 82 70 45 65 68 97 47 19",
		"Card 142: 65 99 24 50 21 68 46 89 29 12 | 12 23 95 96 42 67 82 49 16  8 80  4 85 20 76 89 34 40 44 39 45 86 14 94 61",
		"Card 143: 34 97 20 14  7 31 19 24 16 71 | 16 51 81 20 34  7 46  3 23 83 84 17 73 96 52 31 32 27 24  2 49 44 97 67 14",
		"Card 144:  2 93 72  6 64  3  5 31 56 74 | 64 38 21  6 37 78 20 45 93 58 47 29 55 13 72 14 59 98 25  3 39 99 65 85 31",
		"Card 145: 97 58 24 50 43 86 78 28 73 98 | 38 65 96  3 31 17 74 76 72 48 10 15 85 90 40 75 82 44 87 80 53  1 63 46 67",
		"Card 146: 55 98 17 39 30 14 89 43 91 11 | 63 90 51 73 37 12 91 74 55  3 34 14 89 17 52 84 15 69 44 28 30 81 77 22 47",
		"Card 147: 18 92 20 12 39 30 91 23 69 51 | 51  4 70 33 52 11 50 31 15 27 35 20 93 53 74 84 73 65 43 96 89 79 47 95  8",
		"Card 148: 67 16 96 88 34 92 99 52 58 81 | 71 70 74 45 75 31 25 10 59 17 76 47  8  2 26 28 94 14 12  6 91 20 86 97 11",
		"Card 149: 42 14 48 49 37 60 79 64 59 35 | 89 83  8 46 78  1 13 76 70 82 59 66 84 25 21 65 19 68 72 93 36 11 81 80 90",
		"Card 150: 56 59 61 27 85 80 10 42 43 15 | 52 50 83 65 19 11 69  7 63 38 96 44 30  3 40 58 84 16 88  8 98 90 22 39 62",
		"Card 151: 24  6 35 26 79 71 45 59 98 90 | 36 52 44 84 76 22 21 34 23 94 13 43 49 32  9 57 53 31 29 95 92 88 74 45 61",
		"Card 152: 92 79 98 66 84 64 44 20 32 57 | 83 16 49 41 50  2 72 28 56 81 65 60 54 74 59 26 22 12 70 53 85 73 88 15 35",
		"Card 153: 42 49 39 99 72 12 82 50 29 92 | 20 71 55 78 84 87  2 73 23  8 18 66 31 62 90 26  4 48 54 91 72 80 70 67 17",
		"Card 154: 70 60  1 78  8 62 24  2 67 84 | 55 84 70 38  1 15  6 74 78 62  2 39 71 63 85 11 24 67 90 64 36 46 60 59  8",
		"Card 155:  2 95 93 76 64 89 79 92 35 72 | 79  4 95 30 27 96 98 45 84 23 41 76 34 93 89  2 97 47 72 40 12 64 35 92 11",
		"Card 156: 13 79 80 67 41 46 21 57  4 27 | 73 14 49 57  4 60 90 45 80 21 41 13 67 93 85 79  7 46 91 37 72 32 66 89 27",
		"Card 157: 16 56 29 19 97 91 90 22 78  9 | 74 29 53 99 97 88 85 62 75 38 32 10 51 28 43 14  6 81 39 40 13 17 79 92 35",
		"Card 158: 90 15 84 86 88 10 43 32 78 44 | 10 80 32 88 35 64 23 92 62 83 19 94 57 52 95 67 75 73 21 89 33 74 26  6 55",
		"Card 159: 90 71 30 85 20 12  2 95 79 89 |  8 33 78 81 44 88 82 58 51 40 48 34  7 94 45 84 53 66 17 62 68 27  3 28 96",
		"Card 160: 34 29 24 75 45  7 15 91 88 90 | 91 69 24 90 65 54 38 88 15 72  2  7 53 75 81 89 21 35  3 52 30 83 49 37 34",
		"Card 161: 15 66 33 37  4 19 68 39 86 30 | 65  4 86  6 85 19 26 43 74 50 46  1 22 75 24 27 53 17 14 69 15 73 36 31 30",
		"Card 162: 72 65 87 45 82 88 36  9 57 42 | 85 45 43 47 26  8 72  2 46 65 20 28 76 42 57 90 44 82 50 93 35 78 73 36 88",
		"Card 163:  4 45 52 37  7 93 81 39 64 31 | 60 97 86 37 87 31 76 91 63 40 72 39 52 43 22 64 19 24 82 98 77 26  7 21 83",
		"Card 164: 73 34 98  3 70 23  4 27 86 62 | 34 59  6 62 17 48 33 86 70 53 43 52 89 74 98  4 24 79 50 57 84 88  3 75 67",
		"Card 165: 59  6 18 72 17 53 35 86 51 16 | 72 56 59 40 98 88 15  3 16 84 97 64 96  7 78 80 24  2 86 18 60 51 52 10 85",
		"Card 166: 63 20  3 73 48 60 54 86  2 50 | 54  4 96 55 99 31 73 26 46 89 28 48 27 17 86  7 30 71 33 94 25 66 15 43  3",
		"Card 167: 96 12 93 23 43 49 70 22 76 60 |  4 70  6 19 98  9 26 24 15 85 32 51  8 25 88 68 39 41 97 74 60 71 42  3 94",
		"Card 168: 19 33 79 78 53 47 42 67 14 63 | 15 13 33 54 18 60  6 55 46 79 91 26 58 98  1 65 51 63 27 37 84 96 62 28 10",
		"Card 169: 42 59 85 69 71 24 40 68 67 83 |  6 91 13 57  8 22  4 29 85 94 43 44 55 30 84 98 11 54 71 82 39 90 78  7 80",
		"Card 170: 87 39 98 12 59  5 22 48 78 44 | 69 27 17 64 83 68  5  1 24 43  2 75 89 61 72 95  3 36 38 82 25 10 11 91 87",
		"Card 171: 50 55 88 89 30 59 64 76 49  5 | 28 45 63 53 39 70 36 43 62 79 99 98 95 19  9 97 61 75 15 33 52 10 93 11 90",
		"Card 172: 40 61 73  9 67 22 18 34 76 71 | 14 23  7 11 98 17 97 69 74 48 54 96 56 82 37  1 68  2 57 60 36 51 59 42 10",
		"Card 173: 44 88 19 11 84 91 62 31 73 78 | 72 57 15  9 69 26 11 37 10 56 97  3 51 86 93 25 19 46 87 84 20 31 99 88 61",
		"Card 174: 27  9 36 51 19 64 89 16 53 61 | 51 15 56 84 19 44 96 17 74 64 76  9 61 98 27 36 59 67 32 16 89 70 53 66 24",
		"Card 175: 75 73 27 91 47 15 90 30 48  2 | 27 43 91 95 71 50 90 15 75 58 73  3 48 20  2 31  6  1 89 72 30 66 37 17 68",
		"Card 176:  3 87 49 74 88 50 90  2 10 11 | 68 34 89 67  1  9 64  8 21 27 96 84 48 37 69 28 18 12 97 16 40 93  5 41 77",
		"Card 177: 68  1 51 19 59 52 27 81 78 99 | 91 25 17 35 84 74  9 67 26  2 65 83 98 22 53 18 88 56 63 82 66 39 79 58 96",
		"Card 178: 64 81 17 50 42 57 29 80 16 12 | 54 73 15 89 76 39 42 65 32 16 81  6 29 62 57 46 17 53  7  2 64 44 82 33 50",
		"Card 179: 27 49 28 97 66 31 62 50 20 32 | 31 97 78 94 13 63  1  4 76 28 62 66 57 99 67 74 20  7 27  2 32 49 50 71 45",
		"Card 180: 34 89 40 28 95 80 47 71 31  6 | 80 91 47  6 20 28 85 46 21 18 59 87 83 15 31 40 39  1 34  5 95 71 82 89 88",
		"Card 181: 73 23 70 89 53 63 41  6 10 56 | 54 13 14 90 40 74 29 61 66 36 82 97 21 12 81 15 60 76 52 98 88  2 73 35  1",
		"Card 182: 33 26 55 93 49 35 94 50 45 99 | 93 40 99  3 36 35 49 45 51 82 15 54 88 94 50  8 26  6 84 33 18 59  7 55 29",
		"Card 183: 30 94 37 82 46 36 73 83 16 65 | 94  2 38 44 30 20 36 48 29 26 58 43 33  3 14 39 98 17 76 87 81 86 66 24 63",
		"Card 184: 90 41 81  8 32 88 89 99 44 92 | 41 92 22  3 35 83  8 46 81 44 89 34 88 90 75 69 36 97 32 99 48 53 57 63 38",
		"Card 185: 79  4 70 97 40 23 84 65 28 82 | 99 19 18 55  4 35 49 45 63 82  1 54 59 60 93 67 38 26 47 98 79 92  9 51 66",
		"Card 186: 92 58 39 15 52 86  5 47  9  1 |  1 95 48 37 77 43 69 30 76 96 99 87 34  7 91 66 52 38 16 57 85 62 13 12 51",
		"Card 187: 77 51 68 38 98 76 73 14 41 81 | 41 30 92 65 32 68 38 20 64 54 77 97 14 44 70 91 51 10 57 48 13 76 78 98 40",
		"Card 188: 49 61 74 63 53 22  7 23 87 28 | 55 83 31 91 42 34 48 60 68 72 69 14  8 52  2 30 89 20 92 45 17 51 62  1 70",
		"Card 189: 85 50 70 15 89 93 11 63 60 82 | 67 95 52 98 26 86 35  2 72 43 44 42  6 49  3 11 37 39 51 54  8  4 31 14 41",
		"Card 190: 18 65 32 15 75 48 19 96 69 67 | 62 16 90 58 98 10 78 41 79 25 17 24 20 93 86 49 23 46 13 92 43 34 30 64 51",
		"Card 191: 46 53  3  6 54 19 96 69 44 48 | 78 23 99 38 68  4 35 22 81 55 98 73 43 42 83 80 62 41 66 61 25 89 34 20 54",
		"Card 192: 54 77 76 45 49 99 57 69 86 72 | 91 65 98 79 10  4 57 71 86 80 34 93 73 15  5 59 61 77 88 75  2 48 44 12 56",
		"Card 193: 81 40 62 53  8 10 74 97 12 17 |  7 25 71 38 44 55 51 69 93 40 18 57 32 84 96 17 97 88 34 26 73 86 68 59 35",
		"Card 194: 51 12 21 52 31  6 44 23 76  7 |  1 57 96 55 70 92 46 58 90 50 14 48 49 84 44 74 47 34 87 18 72 56 37 31 17",
		"Card 195: 36 37 17 41 77 88 12 47 94 50 | 83 63 65 35 28 31 38 53 44 13  8 12 69 20 78 48 16 97 21 29 84 94 66 98 75",
		"Card 196: 57 56 99 33 67 55 38 11 81 29 | 17 28  1 98 12 87 69  2 60 96  6 41 89 44 50 84 80 86 54 95  7 52 94  8 30",
		"Card 197: 69 63 15 90 72 56 19 46 27 37 | 96 78 97 50 35 23 62 55  1 71  4 58 57 86 88 22 68 17 83 18 45 94 30 40 14",
		"Card 198: 80 36 74 64 48 57 10 96 31 93 | 87 76 81  9 19 71 13 82 40 31  1 17 84 36 46 26 10 51 94 27 77 38 11 28 35",
		"Card 199: 86 68 30 32 98 92 48 54 65 33 | 31 92 48 86 68 73 15 13 66 32 98  8 82 12 54 65  6 33 83 84 94 56 30 89 19",
		"Card 200: 14 32 69 66 22 60 85 50  6 27 |  1  6 53 86 67 33 56 93 48 16 22 52 85 14 28 91 90  9 17 13 94 20 21 66 50",
		"Card 201: 83 31 70 33 92 24 19 84 61 17 | 10 70 43 98 90 40 31 87 38 50 73 61 18 33 83 54 84 52 20 75 21 19 24 92 17",
		"Card 202: 60 43 55 73 13  5 31 51 40 10 | 14 58 69  7 57 86 53 87 25 98 39 93 92 90 47 66 96 35 44 71 42 77 94 76 65",
		"Card 203: 49 16 55 48 62 61 64 47 32  6 |  1 76 49 67 16 31 99 81 39 15  4 75 55 17 64 35 46 62 87 47 34 58 53  8 21",
		"Card 204: 71 37 46 23  9 28 58 66 75 11 | 58 48 31 66 34 11 72 64 71 15 30 28 46 85  5 37 75  9 23 14 89  2 40 92 47",
		"Card 205: 77 66 32 84 31  7 61 21 62 99 | 77 62 31 36 51 74 21 66 25 13 39 85 54 41 64  8 84 16 86 10  2 99 32 61  7",
		"Card 206: 52 43 11 33 40  7 49 87 25  8 | 77 39 16 25 20 79 65 72 43 26 49 33 52 19 14 60 87 63 85 36  8 38 11 34 47",
		"Card 207: 13 40 42 63 88 35  7 55 81 74 | 77 45  9 83 55 88 37 13 19 81 84 74 46 36 50  7 35 25 11 31 42 10 40 63 73",
		"Card 208: 25 28 12 74 16 13 35 44 53 96 | 43  4 54 42 49 78 87 76 95 35 12  6 44 74 31 28 25 75 30  8 18 53 86 21 39",
		"Card 209: 13  9 49 95 35 78 34  4 33 71 | 51 45 82 59 57 35 34 30 29 66 49 53 60 25 36 40 92 73 77 71 88 81 28 63 43",
		"Card 210: 62 92 40 73  1 45 17 12 85 54 | 17 50  1 43 40 66  2 92 31 34 89 81 11 69 74 10 84 55 79 52 49 83 87 12 96",
		"Card 211: 46 44 30 97 76 83 84 74 26 70 | 28 78 93  6 45 69 90 52 97 44  4 88 46 74 38 91 81 14 63 86 42 39 43 20  7",
		"Card 212: 10 45 38 77 69  4 29 16 59 94 | 59 48 94 55 14 40 92  9  8 41 13 29 37 39 12 16 42 43  1 32 22 63 10 31 51",
		"Card 213: 96 62 42 78 74 88 99 50 20 63 | 74 88 98 67 62  8 33 44 34 82 27  6 13 20 50 43 84 16 59 85 92 87 81 28  9",
		"Card 214: 54 88 80 55 63 69 76 95 86 19 |  7 71 90 74 85 11  3 87 64 78 80 44 20 56 75 59 42 43 96 32 12 77 41 14 81",
		"Card 215:  5 39 47 66 90 45 34 77 80  6 |  7  2 26 76 36  9 88 14 53 93 84 33 50  5 47 69 63 71 66 12 62 39 90  1 51",
		"Card 216:  8 99  6 95 48 33 44 62 26 57 | 97 27 61 71 53 41  4 64 12 60 65 30 58 73 24  3 35 50 11 94 83 19 54 15 36",
		"Card 217: 98 39 72 11 48 76 78 23 18 35 | 73 49 20 17 24 63  9 58 16 44  5 21 96 35 85 19 25 33 43 27 40 52 30 86  4",
		"Card 218: 29 46  2 34 89 12 45  7  8  1 | 14 23 44 67 32 83 41 85 19 33 66 48 77 38 95 50 73 63 29 47 91 15 24  5 60",
		"Card 219: 44 83  7 80 68 17 15  4 45 31 | 41 57 52 79 99 49 98 17 28 82 55 93 50 12 59 62 37 33  1 35 78  6 64 26 43",
		"Card 220: 34 88 44 16 90  6 58 94 64 73 |  5 70 76 53 15 68 28  4 32 65 92 91 24 86 85 31 36 67 83 18 95 45  8 51 74",
	}
	return array
}
