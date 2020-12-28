package code

func minCut(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}
	var dp = make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = i
	}
	for i := 1; i < n; i++ {
		if checkPalindrome(s, 0, i) {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if checkPalindrome(s, j+1, i) {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

func checkPalindrome(str string, left, right int) bool {
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
