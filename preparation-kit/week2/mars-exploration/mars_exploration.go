package marsexploration

import (
	"strings"
)

func MarsExploration(s string) int32 {
	var count int32

	s = strings.ToLower(s)

	arrBadSos := []string{}

	for i := 0; i < len(s)/3; i++ {
		newString := string(s[i*3]) + string(s[(i*3)+1]) + string(s[(i*3)+2])

		if newString != "sos" {
			arrBadSos = append(arrBadSos, newString)
		}
	}

	for i := range arrBadSos {
		if arrBadSos[i][0] != 's' {
			count++
		}

		if arrBadSos[i][1] != 'o' {
			count++
		}

		if arrBadSos[i][2] != 's' {
			count++
		}
	}

	return count
}
