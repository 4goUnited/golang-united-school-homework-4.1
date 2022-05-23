package reverse_string

func ReverseString(input string) string {
	rn := []rune(input)
	for i, j := 0, len(rn)-1; i < len(rn)/2; i, j = i+1, j-1 {
		rn[i], rn[j] = rn[j], rn[i]
	}
	return string(rn)
}
