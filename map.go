package piscine

func Map(f func(int) bool, a []int) []bool {
	arr := make([]bool, len(a))
	if len(a) == 0 {
		return arr
	}
	for i := 0; i < len(a); i++ {
		arr[i] = f(a[i])
	}
	return arr
}
