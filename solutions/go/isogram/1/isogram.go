package isogram

func IsIsogram(word string) bool {
	set := make(map[rune]struct{})
	for _, r := range word {
		if r == ' ' || r == '-' {
			continue
		}

		if r >= 'a' && r <= 'z' {
			r -= 32
		}

		if _, ok := set[r]; ok {
			return false
		}
		set[r] = struct{}{}
	}
	return true
}
