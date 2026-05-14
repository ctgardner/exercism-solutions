package darts

import "math"

func Score(x, y float64) int {
	// Pythagorean theorem: A² + B² = C² ⟹ C = √(A² + B²)
	distance := math.Sqrt(x*x + y*y)

	score := 0
	if distance <= 1 {
		score = 10
	} else if distance <= 5 {
		score = 5
	} else if distance <= 10 {
		score = 1
	}
	return score
}
