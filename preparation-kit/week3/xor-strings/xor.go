package xor

import (
	"strconv"
)

func XOR(a, b string) string {
	intA, _ := strconv.Atoi(a)
	intB, _ := strconv.Atoi(b)

	return strconv.Itoa(intA ^ intB)
}
