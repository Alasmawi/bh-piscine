package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		for i := nb; nb <= nb*nb; i++ {
			if IsPrime(nb) {
				return nb
			}
		}
	}
}
