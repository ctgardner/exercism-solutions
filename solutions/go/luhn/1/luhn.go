package luhn

import "strconv"

func Valid(id string) bool {
	s := []int{}
	for _, r := range id {
		if r == ' ' {
			continue
		}

		i, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}
		s = append(s, i)
	}

	if len(s) < 2 {
		return false
	}

	sum := 0
	for i := len(s) - 1; i >= 0; i -= 2 {
		sum += s[i]
	}
	for i := len(s) - 2; i >= 0; i -= 2 {
		doubled := s[i] * 2
		if doubled > 9 {
			doubled -= 9
		}
		sum += doubled
	}

	return sum%10 == 0
}
