package code

import "math"

func divide(dividend int, divisor int) int {
	signal := -1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		signal = 1
	}
	if dividend > 0 {
		dividend = -dividend
	}
	if divisor > 0 {
		divisor = -divisor
	}

	var result int
	for dividend <= divisor {
		var tmpDivisor = divisor
		var tmpResult = -1
		for dividend <= (tmpDivisor << 1) {
			if tmpDivisor <= (math.MinInt32 >> 1) {
				break
			}
			tmpDivisor = tmpDivisor << 1
			tmpResult = tmpResult << 1
		}
		dividend -= tmpDivisor
		result += tmpResult
	}
	if signal < 0 {
		if result <= math.MinInt32 {
			return math.MaxInt32
		}
		result = -result
	}
	return result
}
