package allyourbase

import (
	"errors"
	"math"
	"slices"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return []int{}, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}

	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	decimal := toDecimal(inputBase, inputDigits)
	return fromDecimal(outputBase, decimal), nil
}

func toDecimal(base int, digits []int) (decimal int) {
	position := 0
	for _, d := range slices.Backward(digits) {
		positionalValue := int(math.Pow(float64(base), float64(position)))
		decimal += d * positionalValue
		position++
	}
	return decimal
}

func fromDecimal(base, decimal int) (digits []int) {
	if decimal == 0 {
		return []int{0}
	}

	for decimal != 0 {
		digits = append(digits, decimal%base)
		decimal /= base
	}
	slices.Reverse(digits)
	return digits
}
