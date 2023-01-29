package main

import (
	"fmt"
)

func main() {
	slice := []int32{1, 2, 3, 4, 3, 2, 1}

	result := lonelyinteger(slice)

	fmt.Println(result)
}

func lonelyinteger(a []int32) int32 {
	// Write your code here

	newMap := make(map[int32]int, len(a))

	for i := range a {
		newMap[a[i]] += 1
	}

	for i := range newMap {
		if newMap[i] == 1 {
			return i
		}
	}

	return 0
}
