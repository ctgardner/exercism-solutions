package raindrops

import "strconv"

var sounds = []struct {
	denominator int
	sound       string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(number int) string {
	sound := ""
	for _, s := range sounds {
		if number%s.denominator == 0 {
			sound += s.sound
		}
	}

	if sound == "" {
		return strconv.Itoa(number)
	} else {
		return sound
	}
}
