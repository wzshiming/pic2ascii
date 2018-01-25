package pic2ascii

func ReverseString(s string) string {
	str := []rune(s)
	l := len(str) / 2
	for i := 0; i < l; i++ {
		j := len(str) - i - 1
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}
