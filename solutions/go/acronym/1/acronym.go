package acronym

import "unicode"

func Abbreviate(s string) string {
	acronym := []rune{}
	prevDelim := true
	for _, r := range s {
		if r == ' ' || r == '-' {
			prevDelim = true
			continue
		}

		if prevDelim {
			r = unicode.ToUpper(r)
			if r >= 'A' && r <= 'Z' {
				acronym = append(acronym, r)
				prevDelim = false
			}
		}
	}

	return string(acronym)
}
