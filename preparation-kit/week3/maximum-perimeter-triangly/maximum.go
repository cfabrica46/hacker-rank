package maximumperimetertriangly

import "sort"

func MaximumPerimeterTriangle(sticks []int32) []int32 {
	sort.Slice(sticks, func(i, j int) bool {
		return sticks[i] > sticks[j]
	})

	for i := 0; i < len(sticks)-2; i++ {
		if sticks[i] < sticks[i+1]+sticks[i+2] {
			return []int32{sticks[i+2], sticks[i+1], sticks[i]}
		}
	}

	return []int32{-1}
}
