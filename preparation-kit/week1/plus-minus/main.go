package main

import "fmt"

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}

	plusMinus(arr)
}

func plusMinus(arr []int32) {
	var nNegative, nZero, nPositive float32

	for i := range arr {
		switch {
		case arr[i] < 0:
			nNegative++
		case arr[i] == 0:
			nZero++
		case arr[i] > 0:
			nPositive++
		}
	}

	fmt.Println(nPositive / float32(len(arr)))
	fmt.Println(nNegative / float32(len(arr)))
	fmt.Println(nZero / float32(len(arr)))
}
