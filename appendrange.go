package piscine

func AppendRange(min, max int) []int {
	arr1 := []int{}

	if min >= max {
		return nil
	}
	for i := min; i < max; i++ {
		arr1 = append(arr1, i)
	}
	return arr1
}
