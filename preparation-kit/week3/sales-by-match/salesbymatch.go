package salesbymatch

func SockMerchant(n int32, ar []int32) int32 {
	var count int32

	myMap := make(map[int32]int32)

	for i := range ar {
		myMap[ar[i]]++
	}

	for i := range myMap {
		if myMap[i] > 1 {
			count += myMap[i] / 2
		}
	}

	return count
}
