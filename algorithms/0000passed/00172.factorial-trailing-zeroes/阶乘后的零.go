package code

func trailingZeroes(n int) int {
	var zeroCount int
	for n > 0 {
		n /= 5
		zeroCount += n
	}
	return zeroCount
}
