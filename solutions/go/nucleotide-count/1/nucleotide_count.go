package nucleotidecount

import "errors"

type Histogram map[rune]int
type DNA []rune

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, r := range d {
		if r != 'A' && r != 'C' && r != 'G' && r != 'T' {
			return nil, errors.New("invalid character")
		}
		h[r]++
	}
	return h, nil
}
