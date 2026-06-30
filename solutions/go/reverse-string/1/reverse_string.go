package reversestring

import "unicode/utf8"

func Reverse(input string) string {
	l := utf8.RuneCountInString(input)
	rs := make([]rune, l)
	for i, r := range []rune(input) {
		rs[l-1-i] = r
	}
	return string(rs)
}
