package rnatranscription

var complements = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	rna := []rune(dna)
	for i, r := range rna {
		rna[i] = complements[r]
	}
	return string(rna)
}
