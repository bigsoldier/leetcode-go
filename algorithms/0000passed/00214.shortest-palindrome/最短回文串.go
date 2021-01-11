package code

func shortestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	i := getIndex(s + "#" + reverse(s))
	return reverse(s[i:]) + s
}

// 反转字符串
func reverse(s string) string {
	n := len(s)
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = s[n-i-1]
	}
	return string(bytes)
}

func getIndex(pattern string) int {
	n := len(pattern)
	next := make([]int, n)
	k, i := 0, 1
	for i < n {
		if pattern[i] == pattern[k] {
			next[i] = k + 1
			k = next[i]
			i++
		} else if k == 0 {
			next[i] = k
			i++
		} else {
			k = next[k-1]
		}
	}
	return next[n-1]
}
