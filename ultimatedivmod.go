package piscine

func UltimateDivMod(a *int, b *int) {
	*a /= *b
	*b = *a % *b
}
