package diffsquares

func SquareOfSum(n int) int {
	res := (n + 1) * n / 2
	return res * res
}

func SumOfSquares(n int) int {
	res := n * (n + 1) * (2 * n + 1) / 6
	return res
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
