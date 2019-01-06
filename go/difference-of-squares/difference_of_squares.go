// Package diffsquares contains utilities for calculating the sum of squares,
// the square of sums, and the difference between them
package diffsquares

// SquareOfSum returns the square of sums for all numbers up to the given number
func SquareOfSum(num int) int {
	sum := (num * (num + 1)) / 2
	return sum
}

// SumOfSquares returns the sum of squares for all numbers up to the given number
func SumOfSquares(num int) int {
	return (num * (num + 1) * (2*num + 1)) / 6
}

// Difference returns the difference between the square of the sum and the sum
// of the squares for all numbers up to the given number
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
