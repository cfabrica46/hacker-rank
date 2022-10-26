package main

import "fmt"

func main() {
	arr := []int32{256741038, 623958417, 467905213, 714532089, 938071625}

	minMaxSum(arr)
}

func minMaxSum(arr []int32) {
	quicksortRecursive(arr, 0, int32(len(arr)-1))

	newArr := make([]int64, len(arr))

	for i := range arr {
		newArr[i] = int64(arr[i])
	}

	fmt.Println(newArr[0]+newArr[1]+newArr[2]+newArr[3], newArr[1]+newArr[2]+newArr[3]+newArr[4])
}

func quicksortRecursive(slice []int32, lo, hi int32) {
	if lo > hi {
		return
	}

	indexPivot := partitionRecursive(slice, lo, hi)

	quicksortRecursive(slice, lo, indexPivot-1)
	quicksortRecursive(slice, indexPivot+1, hi)
}

func partitionRecursive(slice []int32, lo, hi int32) int32 {
	pivot := slice[hi]

	for i := lo; i < hi; i++ {
		if slice[i] < pivot {
			slice[i], slice[lo] = slice[lo], slice[i]
			lo++
		}
	}

	slice[lo], slice[hi] = slice[hi], slice[lo]

	return lo
}
