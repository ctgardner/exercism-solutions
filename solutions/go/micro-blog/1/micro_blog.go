package microblog

func Truncate(phrase string) string {
	s := []rune(phrase)
	if len(s) < 5 {
		return string(s)
	}
	return string(s[:5])
}
