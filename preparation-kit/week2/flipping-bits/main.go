package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := int64(9)

	result := flippingBits(n)

	fmt.Println(result)
}

func flippingBits(n int64) int64 {
	binary := convertDecimalToBinary(n)

	fmt.Println(binary)

	binaryString := strconv.FormatUint(binary, 10)

	binaryString = completeZeros(binaryString)

	binaryString = inverseBinary(binaryString)

	return convertBinaryToDecimal(binaryString)
}

func convertDecimalToBinary(n int64) uint64 {
	binary := uint64(0)
	base := uint64(1)

	for n > 0 {
		remainder := n % 2
		binary += uint64(remainder) * base
		n = n / 2
		base *= 10
	}

	return binary
}

func completeZeros(binaryString string) string {
	missingsZeros := 32 - len(binaryString)

	var zeros string

	for i := 0; i < missingsZeros; i++ {
		zeros += "0"
	}

	return zeros + binaryString
}

func inverseBinary(binaryString string) (newBinaryString string) {
	for i := range binaryString {
		if binaryString[i] == []byte("0")[0] {
			newBinaryString = newBinaryString + "1"
		} else {
			newBinaryString = newBinaryString + "0"
		}
	}

	return newBinaryString
}

func convertBinaryToDecimal(binary string) int64 {
	decimal := 0
	base := 1

	for i := len(binary) - 1; i >= 0; i-- {
		if binary[i] == '1' {
			decimal += base
		}
		base *= 2
	}

	return int64(decimal)
}
