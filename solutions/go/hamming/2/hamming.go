package hamming

import "errors"

func Distance(a, b string) (int, error) {
	// We can safely iterate and compare bytes (instead of runes)
	// because all characters in a DNA strand are ASCII. ie, ACGT
	if len(a) != len(b) {
		return 0, errors.New("arguments must have same length")
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
