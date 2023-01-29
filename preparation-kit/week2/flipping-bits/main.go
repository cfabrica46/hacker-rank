package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := int64(802743475)

	result := flippingBits(n)

	fmt.Println(result)
}

func flippingBits(n int64) int64 {
	binary := convertDecimalToBinary(n)

	binary = completeZeros(binary)

	binary = inverseBinary(binary)

	return convertBinaryToDecimal(binary)
}

func convertDecimalToBinary(decimal int64) string {
	bin := strconv.FormatInt(decimal, 2)

	return bin
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

func combineStrings(strgs []string) (s string) {
	for i := range strgs {
		s += strgs[i]
	}

	return s
}
