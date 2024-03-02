package piscine

func AppendRange(min, max int) []int {
	arr1 := []int{min}

	if min >= max || (min < 0 || max < 0) {
		return arr1
	}
	index := 0
	for i := min + 1; i < max; i++ {
		arr1[index] = i
		index++
	}
	return arr1
}
