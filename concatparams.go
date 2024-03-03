package piscine

func ConcatParams(args []string) string {
	str := ""
	for i := 0; i < len(args); i++ {
		if !(i == len(args)-1) {
			str += args[i] + "\n"
		} else {
			str += args[i]
		}
	}
	return str
}
