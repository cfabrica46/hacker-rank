package marsexploration

import (
	"strings"
)

func MarsExploration(s string) int32 {
	var count int32

	s = strings.ToLower(s)

	for i := 0; i < len(s); i += 3 {
		if s[i] != 's' {
			count++
		}

		if s[i+1] != 'o' {
			count++
		}

		if s[i+2] != 's' {
			count++
		}
	}

	return count
}
