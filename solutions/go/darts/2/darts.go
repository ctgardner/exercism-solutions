package darts

import "math"

func Score(x, y float64) int {
	// Pythagorean theorem: A² + B² = C² ⟹ C = √(A² + B²)
	distance := math.Sqrt(x*x + y*y)

	var score int
	switch {
	case distance <= 1:
		score = 10
	case distance <= 5:
		score = 5
	case distance <= 10:
		score = 1
	default:
		score = 0
	}
	return score
}
