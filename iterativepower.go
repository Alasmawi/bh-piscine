package piscine

func IterativePower(nb int, power int) int {
	if power < 0 || power > 20 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		nnb := nb

		for i := 0; i < power; i++ {
			nnb *= nb
		}
		return nnb

	}
}
