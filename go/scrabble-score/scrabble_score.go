package scrabblescore

import "unicode"

func Score(word string) int {
	score := 0
	for _, r := range word {
		value := 0
		switch unicode.ToUpper(r) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			value = 1
		case 'D', 'G':
			value = 2
		case 'B', 'C', 'M', 'P':
			value = 3
		case 'F', 'H', 'V', 'W', 'Y':
			value = 4
		case 'K':
			value = 5
		case 'J', 'X':
			value = 8
		case 'Q', 'Z':
			value = 10
		}
		score += value
	}
	return score
}
