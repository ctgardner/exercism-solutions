package isogram

func IsIsogram(word string) bool {
	var bitfield uint32 = 0
	for _, r := range word {
		if r == ' ' || r == '-' {
			continue
		}

		if r >= 'a' && r <= 'z' {
			r -= 32
		}

		mask := uint32(1 << (r - 'A'))
		if bitfield & mask != 0 {
			return false
		}
		bitfield |= mask
	}
	return true
}
