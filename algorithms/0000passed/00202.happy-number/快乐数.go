package code

func isHappy(n int) bool {
	m := map[int]bool{}
	for n != 1 && !m[n] {
		n, m[n] = step(n), true
	}
	return n == 1
}

func step(n int) int {
	var sum int
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
