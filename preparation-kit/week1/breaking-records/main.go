package main

import "fmt"

func main() {
	scores := []int32{12, 24, 10, 24}

	result := breakingRecords(scores)

	fmt.Println(result)
}

func breakingRecords(scores []int32) (breakRecords []int32) {
	// Write your code here
	breakRecords = []int32{0, 0}

	max := scores[0]
	min := scores[0]

	for i := 1; i < len(scores); i++ {
		if scores[i] > max {
			max = scores[i]
			breakRecords[0]++

		} else if scores[i] < min {
			min = scores[i]
			breakRecords[1]++
		}
		fmt.Println(breakRecords)
	}

	return breakRecords
}
