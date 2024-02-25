package piscine

func IsPrime(nb int) bool {
	count := 0
	if nb <= 1 {
		return false
	} else if nb%2 == 0 && nb != 2 {
		return false
	} else {
		for i := 1; i <= nb; i++ {
			if nb%i == 0 {
				count++
			}
		}
	}
	if count == 2 {
		return true
	} else {
		return false
	}
}
