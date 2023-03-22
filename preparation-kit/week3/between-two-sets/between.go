package betweentwosets

import (
	"fmt"
	"sort"
)
//test

/* func GetTotalX(a []int32, b []int32) int32 {
	// Write your code here

	var count int32

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })

	for i := a[len(a)-1]; i <= b[0]; i++ {
		check := true

		for j := range a {
			if i%a[j] != 0 {
				check = false
			}
		}

		for j := range b {
			if b[j]%i != 0 {
				check = false
			}
		}

		if check {
			count++
		}
	}

	return count
} */

// Función para encontrar el máximo común divisor de dos números
func gcd(a, b int32) int32 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Función para encontrar el mínimo común múltiplo de un slice de números
func lcm(numbers []int32) int32 {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = result * numbers[i] / gcd(result, numbers[i])
	}
	return result
}

// Función para verificar si un número es un factor común de los elementos de B
func isFactorOfB(num int32, arrB []int32) bool {
	for _, b := range arrB {
		if b%num != 0 {
			return false
		}
	}
	return true
}

// Función para encontrar todos los números enteros que satisfacen las condiciones
func GetTotalX(arrA []int32, arrB []int32) int32 {
	// Ordenamos el slice B en orden ascendente
	sort.Slice(arrB, func(i, j int) bool { return arrB[i] < arrB[j] })

	// Encontramos el mínimo común múltiplo de A
	lcmA := lcm(arrA)

	// fmt.Println(lcmA)

	count := int32(0)
	// Probamos todos los múltiplos del mínimo común múltiplo de A hasta el primer elemento de B
	for num := lcmA; num <= arrB[0]; num += lcmA {
		if isFactorOfB(num, arrB) {
			fmt.Println(num)
			count++
		}
	}
	return count
}
