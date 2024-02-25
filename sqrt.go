package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 0
	} else if nb == 1 {
		return 1
	} else {
		sqrt := 0
		for i := 0; i*i <= nb; i++ {
			if i*i == nb {
				sqrt = i
			}
		}
		return sqrt

	}
}
