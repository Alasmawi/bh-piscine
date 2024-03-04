package piscine

func Map(f func(int) bool, a []int) []bool {
	arr := []bool{}
	for i := 0; i < len(a); i++ {
		arr[i] = f(a[i])
	}
	return arr
}
