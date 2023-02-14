package permutingtwoarrays

import (
	"fmt"
	"sort"
)

func TwoArrays(k int32, A []int32, B []int32) string {
	n := len(A)

	// Se ordena de esta manera para que las sumas sean mayores a k
	// de manera que el valor mas pequeño de A se suma con el mas grande
	// de B, si estuvieran ambos ordenados de la misma manera se podria
	// sumar el numero mas pequeño de ambos dificultando el superar a k

	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })
	sort.Slice(B, func(i, j int) bool { return B[i] > B[j] })

	for i := 0; i < n; i++ {
		if A[i]+B[i] < k {
			return "NO"
		}
	}
	return "YES"
}

// ---

func MyTwoArrays(k int32, a []int32, b []int32) string {
	nA := 0
	nB := 0

	for {
		for i := range a {
			fmt.Printf("A = %d | B = %d, A+B = %d | K = %d\n", a[i], b[i], a[i]+b[i], k)
			if a[i]+b[i] < k {
				break
			} else {
				if i == len(a)-1 {
					return "YES"
				}
			}
		}

		if nA < len(a) {
			a = permutArray(a)
			nA++
		} else if nB < len(b) {
			b = permutArray(b)
			nB++
		} else {
			return "NO"
		}
	}
}

func permutArray(array []int32) []int32 {
	n := array[0]

	newArr := array[1:]

	newArr = append(newArr, n)

	return newArr
}
