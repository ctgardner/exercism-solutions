package raindrops

import "strconv"

func Convert(number int) string {
	if number % 3 != 0 && number % 5 != 0 && number % 7 != 0 {
		return strconv.Itoa(number)
	}

	sound := ""
	if number % 3 == 0 {
		sound += "Pling"
	}
	if number % 5 == 0 {
		sound += "Plang"
	}
	if number % 7 == 0 {
		sound += "Plong"
	}
	return sound
}
