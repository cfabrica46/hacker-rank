package pangrams

import (
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func Pangrams(s string) string {
	if len(s) < len(alphabet) {
		return "not pangram"
	}

	s = strings.ToLower(s)

	letter := make(map[rune]bool)

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			letter[v] = true
		}
	}

	for c := 'a'; c <= 'z'; c++ {
		if !letter[c] {
			return "not pangram"
		}
	}

	return "pangram"
}
