package piscine

func BasicAtoi(s string) int {
result:=0
	for_, r:= range s{

		digit:= int(r-'0')
		result=result*10+digit
	}
	return result
}
