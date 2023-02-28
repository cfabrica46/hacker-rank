package migratorybirds

import (
	"fmt"
)

func MigratoryBirds(arr []int32) (result int32) {
	myMap := make(map[int32]int32)

	for i := range arr {
		myMap[arr[i]]++
	}

	var maxValue int32
	var maxID int32

	for i := range myMap {
		if myMap[i] >= maxValue {
			fmt.Println(myMap[i], maxValue, i, maxID)

			if i < maxID {
				maxID = i
				maxValue = myMap[i]
			}
			fmt.Println("->", myMap[i], maxValue, i, maxID)
		}
	}

	return maxID
}
