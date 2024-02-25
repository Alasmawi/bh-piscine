package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else {
		nnb := nb

		for i := 1; i < power; i++ {
			nnb = nnb * nb
		}

		return nnb

	}
}
