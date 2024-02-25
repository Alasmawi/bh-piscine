package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		for i := nb + 1; nb <= nb*nb; i++ {
			if IsPrime(nb) {
				return nb
			}
		}
	}
	return 0
}
