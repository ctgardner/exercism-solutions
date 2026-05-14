package darts

import "math"

var scoringCircles = []struct {
	radius float64
	score  int
}{
	{1.0, 10},
	{5.0, 5},
	{10.0, 1},
}

func Score(x, y float64) int {
	// Pythagorean theorem: A² + B² = C² ⟹ C = √(A² + B²)
	distance := math.Sqrt(x*x + y*y)

	for _, c := range scoringCircles {
		if distance <= c.radius {
			return c.score
		}
	}
	return 0
}
