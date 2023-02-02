package main

import (
	"fmt"
)

func main() {
	res := countingValleys(12, "DDUUDDUDUUUD")

	fmt.Println(res)
}

func countingValleys(_ int32, path string) int32 {
	var valleys int32
	level := 0

	// path = strings.ToUpper(path)

	for i := range path {
		if path[i] == 'U' {
			level++

			if level == 0 {
				valleys++
			}

		} else {
			level--
		}
	}

	return valleys
}
