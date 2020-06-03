package code

func grayCode(n int) []int {
	var gray = []int{0}
	var head = 1
	for i := 0; i < n; i++ {
		for j := len(gray) - 1; j >= 0; j-- {
			gray = append(gray, head+gray[j])
		}
		head <<= 1
	}
	return gray
}
