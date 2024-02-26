package piscine

func FindNextPrime(nb int) int {
	for i := nb + 1; nb <= nb*nb; i++ {
		if IsPrime(nb) {
			return nb
		}
		nb++
	}
}
