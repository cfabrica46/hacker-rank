package main

import "fmt"

func main() {
	arr := []int32{1, 2, 3, 4}

	fmt.Println(arr)
	fmt.Println(reverseArray(arr))
}

func reverseArray(a []int32) []int32 {
	var newArr []int32

	for i := len(a) - 1; i >= 0; i-- {
		newArr = append(newArr, a[i])
	}

	return newArr
}
