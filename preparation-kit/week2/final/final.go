package final

import "fmt"

func FlippingMatrix(matrix [][]int32) int32 {
	n := len(matrix) / 2
	var max int32

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			p1 := matrix[i][j]
			p2 := matrix[2*n-1-i][j]
			p3 := matrix[i][2*n-1-j]
			p4 := matrix[2*n-1-i][2*n-1-j]

			max += int32(maxToInts(p1, p2, p3, p4))

			fmt.Println(p1, p2, p3, p4, max)
		}
	}

	return max
}

func maxToInts(a ...int32) int32 {
	m := a[0]
	for _, x := range a[1:] {
		if x > m {
			m = x
		}
	}

	fmt.Println("Max:", m)
	return m
}
