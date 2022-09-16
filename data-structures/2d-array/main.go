package main

import "fmt"

func main() {
	arr := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		//---
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}

	fmt.Println(hourglassSum(arr))
}

func hourglassSum(arr [][]int32) int32 {
	var max int32

	fmt.Println("----Start----")
	/* var height, weight int

	height = len(arr)
	weight = len(arr[0])

	fmt.Println(height, weight) */

	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr[0])-2; j++ {
			auxSum := sum(arr, i, j)

			if max < auxSum {
				max = auxSum
			}
			/* fmt.Println("*******")
			fmt.Println(arr[i][j], arr[i][j+1], arr[i][j+2])
			fmt.Println(arr[i+1][j+1])
			fmt.Println(arr[i+2][j], arr[i+2][j+1], arr[i+2][j+2])
			fmt.Println("*******") */
		}
		// fmt.Println("")
	}

	fmt.Println("----END----")
	return max
}

func sum(arr [][]int32, i, j int) int32 {
	return arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
}
