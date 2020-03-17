package code

import "math"

func reverse(x int) int {
	var temp int
	for ; x != 0; x /= 10 {
		t := x % 10
		if temp > math.MaxInt32/10 || (temp == math.MaxInt32/10 && t > 7) {
			return 0
		}
		if temp < math.MinInt32/10 || (temp == math.MinInt32/10 && t < -8) {
			return 0
		}
		temp = temp*10 + t
	}
	return temp
}
