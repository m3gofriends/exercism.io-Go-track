// Package diffsquares implements a simple mathematical calculation.
package diffsquares

// SquareOfSum calculates the square of the sum.
func SquareOfSum(n int) int {
	sum := (1 + n) * n / 2
	return sum * sum
}

// SumOfSquares calculates the sum of the square.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns the difference between SquareOfSum minus SumOfSquares.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
