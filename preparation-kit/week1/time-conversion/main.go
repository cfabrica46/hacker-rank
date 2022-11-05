package main

import (
	"fmt"
	"time"
)

func main() {
	time := "03:04:05PM"

	fmt.Println(timeConversion(time))
}

func timeConversion(s string) string {
	layout1 := "03:04:05PM"
	layout2 := "15:04:05"

	t, _ := time.Parse(layout1, s)

	return t.Format(layout2)
}
