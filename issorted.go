package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	bol := true
	bol1 := true
	for i := 0; i < len(a)-1; i++ {
		if (f(a[i+1], a[i])) < 0 {
			bol = false
		}
	}
	for i := len(a) - 1; i >= 0; i-- {
		if (f(a[i-1], a[i])) < 0 {
			bol = false
		}
	}
	return bol || bol1
}
