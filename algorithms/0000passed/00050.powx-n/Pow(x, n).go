package code

func myPow(x float64, n int) float64 {
	var ret float64 = 1
	if n < 0 {
		x = 1 / x
		n = -n
	}

	var current = x
	for i := n; i > 0; i /= 2 {
		if i%2 == 1 {
			ret *= current
		}
		current *= current
	}
	return ret
}
