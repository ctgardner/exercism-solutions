package differenceofsquares

func SquareOfSum(n int) int {
	// https://en.wikipedia.org/wiki/Squared_triangular_number
	sum := (n * (n + 1)) / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	// https://en.wikipedia.org/wiki/Square_pyramidal_number
	return (n * (n + 1) * (2*n + 1)) / 6
}

func Difference(n int) int {
	// Guaranteed to be non-negative, for N > 0
	return SquareOfSum(n) - SumOfSquares(n)
}
