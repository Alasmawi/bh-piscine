package piscine

func Split(s, sep string) []string {
	index := 0
	var o string
	var wee []string
	for i := 0; i < len(s)-len(sep)+1; i++ {
		if s[i:i+len(sep)] == sep {
			o = s[index:i]
			wee = append(wee, o)
			o = ""
			index = i + len(sep)
		}
	}
	// Append the remaining substring after the last occurrence of the separator
	o = s[index:]
	wee = append(wee, o)
	return wee
}
