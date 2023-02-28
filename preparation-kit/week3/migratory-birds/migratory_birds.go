package migratorybirds

func MigratoryBirds(arr []int32) int32 {
	myMap := make(map[int32]int32)

	for i := range arr {
		myMap[arr[i]]++
	}

	var maxValue int32
	var maxID int32

	for i := range myMap {
		if myMap[i] >= maxValue {
			if myMap[i] > maxValue || i < maxID {
				maxID = i
				maxValue = myMap[i]
			}
		}
	}

	return maxID
}
