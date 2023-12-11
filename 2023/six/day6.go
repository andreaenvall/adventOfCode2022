package main

import "fmt"

func main() {
	puzzle := [][]int{
		{38, 241},
		{94, 1549},
		{79, 1074},
		{70, 1091},
	}
	allrecords := []int{}
	for _, race := range puzzle {
		record := findTimesToBeatRecord(race[0], race[1])
		allrecords = append(allrecords, record)
	}
	result := 1
	for _, v := range allrecords {
		result *= v
	}
	fmt.Println(result)

	result1 := findTimesToBeatRecord(38947970, 241154910741091)
	fmt.Println(result1)

}

func findTimesToBeatRecord(time, record int) int {
	holdBtn := 0
	distance := 0
	higherThanRecord := []int{}
	for i := 0; i < time; i++ {
		holdBtn = i
		distance = holdBtn * (time - holdBtn)
		if distance > record {
			higherThanRecord = append(higherThanRecord, distance)
		}
	}
	return len(higherThanRecord)
}
