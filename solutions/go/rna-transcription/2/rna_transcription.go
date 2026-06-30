package rnatranscription

import "strings"

var complements = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	return strings.Map(func(r rune) rune {
		return complements[r]
	}, dna)
}
