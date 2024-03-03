package piscine

func AppendRange(min, max int) []int {
	arr1 := []int{}

	if min >= max || (min < 0 || max < 0) {
		return arr1
	}
	for i := min; i < max; i++ {
		arr1 = append(arr1, i)
	}
	return arr1
}
