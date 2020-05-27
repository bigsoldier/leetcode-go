package code

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	// 检查字符是否相等
	n := len(s1)
	var chars = make([]int, 256)
	for i := 0; i < n; i++ {
		chars[s1[i]]++
		chars[s2[i]]--
	}
	for _, char := range chars {
		if char != 0 {
			return false
		}
	}

	// 检查字符串是否scramble
	for i := 1; i < n; i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i]) {
			return true
		}
	}
	return false
}
