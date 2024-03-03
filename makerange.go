package piscine

func MakeRange(min, max int) []int {
	size := max - min
	if size < 0 {
		size *= -1
	}
	arr1 := make([]int, size)
	if min >= max {
		return nil
	}
	index := 0
	for i := min; i < max; i++ {
		arr1[index] = i
		index++
	}
	return arr1
}
