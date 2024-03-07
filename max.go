package piscine

func Max(a []int) int {
	temp := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > temp {
			temp = a[i]
		}
	}
	return temp
}
