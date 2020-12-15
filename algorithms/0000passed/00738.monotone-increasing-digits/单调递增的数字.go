package code

import "strconv"

func monotoneIncreasingDigits(N int) int {
	s := []byte(strconv.Itoa(N))
	i := 1
	for i < len(s) && s[i-1] <= s[i] {
		i++
	}

	if i < len(s) {
		for i > 0 && s[i-1] > s[i] {
			s[i-1]--
			i--
		}
		for i++; i < len(s); i++ {
			s[i] = '9'
		}
	}
	ret, _ := strconv.Atoi(string(s))
	return ret
}
