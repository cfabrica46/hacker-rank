package pickingnumbers

import "fmt"

// PickingNumbers encuentra la subsecuencia más larga de enteros dentro de a donde la diferencia
// entre cualquier par de números adyacentes no sea mayor que 1
/* func PickingNumbers(a []int32) int32 {
	var maxLength int32
	lastStop := 0

	for i := 0; i < len(a)-1; i++ {
		if a[i]-a[i+1] > 1 || a[i]-a[i+1] < -1 {

			currentLength := i + 1 - lastStop

			if maxLength < int32(currentLength) {
				maxLength = int32(currentLength)
			}

			lastStop = i + 1
		}
	}

	currentLength := int32(len(a) - lastStop)

	if maxLength < currentLength {
		maxLength = currentLength
	}

	return maxLength
} */

// PickingNumbers encuentra la longitud máxima de una subsecuencia en a donde la diferencia entre
// cada par de elementos adyacentes es como máximo 1
func PickingNumbers(a []int32) int32 {
	freq := make(map[int32]int32)
	var maxLength int32

	for _, num := range a {
		freq[num]++
	}

	fmt.Println(freq)

	for num := range freq {
		fmt.Println("num:", num)
		currLength := freq[num] + freq[num+1]
		fmt.Println("currLength:", currLength)
		if currLength > maxLength {
			maxLength = currLength
		}
	}

	return maxLength
}
