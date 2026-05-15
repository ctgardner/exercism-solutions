package hamming

import "errors"

func Distance(a, b string) (int, error) {
	ra, rb := []rune(a), []rune(b)
	if len(ra) != len(rb) {
		return 0, errors.New("arguments must have same length")
	}

	distance := 0
	for i := 0; i < len(ra); i++ {
		if ra[i] != rb[i] {
			distance++
		}
	}
	return distance, nil
}
