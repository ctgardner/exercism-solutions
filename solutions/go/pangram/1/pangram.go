package pangram

func IsPangram(input string) bool {
	set := make(map[rune]struct{})
	for _, r := range input {
		if r >= 'a' && r <= 'z' {
			r -= 'a' - 'A' // convert lowercase to uppercase
		} else if r < 'A' || r > 'Z' {
			continue // ignore non-alphabetical characters
		}
		set[r] = struct{}{}
	}
	return len(set) == 26
}
