package differenceofsquares

func SquareOfSum(n int) int {
	v := 0
    for i := 1; i <= n; i++ {
        v += i
    }
    return v*v
}

func SumOfSquares(n int) int {
	v := 0
    for i := 1; i <= n; i++ {
        v += i*i
    }
    return v
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
