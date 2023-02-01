package main

import (
	"fmt"
)

func main() {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	res := diagonalDifference(arr)

	fmt.Println(res)
}

func diagonalDifference(arr [][]int32) int32 {
	var n1, n2 int32

	for i := range arr {
		n1 += arr[i][i]
		n2 += arr[i][len(arr)-1-i]
	}

	if n1 > n2 {
		return n1 - n2
	}

	return n2 - n1
}
