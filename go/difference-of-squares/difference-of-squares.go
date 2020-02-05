/*
Package diffsquares calculates the difference between the sum of squares and square of sums. Includes tests.
*/
package diffsquares

//SquareOfSum Calculates the square of the sum up to n
func SquareOfSum(n int) (square int) {
	s := (n*n + n) / 2
	square = s * s
	return square
}

//SumOfSquares Calculates the sum of the squares up to n
func SumOfSquares(n int) (sum int) {
	sum = (2*n*n*n + 3*n*n + n) / 6
	return sum
}

//Difference Calculates the difference between the square of the sum and sum of the squares up to n.
func Difference(n int) (diff int) {
	diff = SquareOfSum(n) - SumOfSquares(n)
	return diff
}
