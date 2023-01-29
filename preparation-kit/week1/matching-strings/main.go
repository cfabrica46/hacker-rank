package main

import "fmt"

func main() {
	strings := []string{"aba", "baba", "aba", "xzxb"}
	queries := []string{"aba", "xzxb", "ab"}

	res := matchingStrings(strings, queries)

	fmt.Println(res)
}

func matchingStrings(strings []string, queries []string) []int32 {
	// Write your code here
	matches := make([]int32, len(queries))

	for i := range queries {
		for j := range strings {
			if queries[i] == strings[j] {
				matches[i]++
			}
		}
	}

	return matches
}
