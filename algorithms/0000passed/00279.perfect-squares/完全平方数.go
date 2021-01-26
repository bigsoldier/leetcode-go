package code

import "math"

func numSquares(n int) (cnt int) {
	for n&3 == 0 {
		n >>= 2 // 4^k
	}
	if n&7 == 7 { // mod 8
		return 4
	}
	if isSquare(n) {
		return 1
	}
	for i := 1; i <= int(math.Pow(float64(n), 0.5)+1); i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}
	return 3
}

func isSquare(n int) bool {
	sq := math.Sqrt(float64(n))
	return int(sq)*int(sq) == n
}
