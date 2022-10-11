package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		//---
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}

	fmt.Println(hourglassSum(arr))
}

func hourglassSum(arr [][]int32) int32 {
	// var max int32
	var arrSum []int

	// fmt.Println("----Start----")

	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr[0])-2; j++ {
			auxSum := sum(arr, i, j)

			/*
				Fails because max init with 0 value and auxSum is minor of 0 for example = -6
				needs other solution

				if max < auxSum {
					max = auxSum
				}
			*/

			arrSum = append(arrSum, int(auxSum))
			sort.Ints(arrSum)

		}
	}

	// fmt.Println("----END----")
	return int32(arrSum[len(arrSum)-1])
}

func sum(arr [][]int32, i, j int) int32 {
	return arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
}
