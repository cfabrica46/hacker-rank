package subrraydivision

func Birthday(s []int32, d, m int32) int32 {
	var count int32

	aum := int(m)

	for i := 0; i < len(s)-(aum-1); i++ {
		res := sumAll(s[i : aum+i]...)

		if res == d {
			count++
		}
	}

	return count
}

func sumAll(n ...int32) (sum int32) {
	for i := range n {
		sum += n[i]
	}

	return sum
}
