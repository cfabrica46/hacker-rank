package main

import "fmt"

func main() {
	grades := []int32{73, 29, 57}

	result := gradingStudents(grades)

	fmt.Println(result)
}

func gradingStudents(grades []int32) []int32 {
	newArr := make([]int32, len(grades))

	for i := range grades {
		newArr[i] = grades[i]

		if grades[i] >= 38 {
			if grades[i]%5 >= 3 {
				newArr[i] = grades[i] - (grades[i] % 5) + 5
			}
		}

	}

	return newArr
}
