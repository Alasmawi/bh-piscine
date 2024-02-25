package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 24 {
		return 0
	}
	factorial := 1
	for i := 1; i <= nb; i++ {
		factorial = factorial * i
	}

	return factorial
}
