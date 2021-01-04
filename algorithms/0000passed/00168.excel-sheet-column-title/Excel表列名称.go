package code

func convertToTitle(n int) (ans string) {
	for n > 0 {
		n--
		ans = string('A'+n%26) + ans
		n = n / 26
	}
	return
}
