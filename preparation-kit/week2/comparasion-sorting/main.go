package main

import "fmt"

func main() {
	arr := []int32{5, 6, 1, 3, 10, 12, 1}

	res := countingSort(arr)

	fmt.Println(res)
}

func countingSort(arr []int32) []int32 {
	var newArr [100]int32

	for i := range arr {
		newArr[arr[i]]++
	}

	arr = make([]int32, len(arr))

	index := 0

	for i := range newArr {
		for j := 0; j < int(newArr[i]); j++ {
			arr[index] = int32(i)
			index++
		}
	}

	return arr
}
