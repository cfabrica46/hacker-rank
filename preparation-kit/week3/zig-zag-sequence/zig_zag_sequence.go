package zigzagsequence

import (
	"fmt"
	"sort"
)

func FindZigZagSequence(arr []int32, n int) []int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	mid := (n - 1) / 2
	arr[mid], arr[n-1] = arr[n-1], arr[mid]

	fmt.Println(arr)

	st := mid + 1
	ed := n - 2

	fmt.Println(st, ed)

	for st <= ed {
		arr[st], arr[ed] = arr[ed], arr[st]
		st++
		ed--
	}

	return arr
}
