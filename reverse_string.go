package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)
	stringLength := len(runes)
	for i := 0; i < stringLength/2; i++ {
		runes[i], runes[stringLength-i-1] = runes[stringLength-i-1], runes[i]
	}
	return string(runes)
}
