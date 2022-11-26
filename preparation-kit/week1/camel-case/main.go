package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var cases []string

	for i := 0; i < 4; i++ {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		cases = append(cases, text)
	}

	fmt.Println("---")

	for i := range cases {
		s := strings.Split(cases[i], ";")

		if s[0] == "S" {
			switch s[1] {
			case "V":
				/* re := regexp.MustCompile(`(?<!^)(?=[A-Z])`)

				result := re.FindAll([]byte(s[2]), -1)

				fmt.Printf("%q\n", result) */
			}
		} else {
			switch s[1] {
			case "M":
				result := strings.ToUpper(string(s[2][0]))
				fmt.Println(result)
				result += s[2][1:]
				fmt.Println(result)
			}
		}
	}
}
