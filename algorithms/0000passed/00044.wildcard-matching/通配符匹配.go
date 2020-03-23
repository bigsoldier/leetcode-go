package code

func isMatch(s string, p string) bool {
	var sLen, pLen = len(s), len(p)
	var dp = make([][]bool, sLen+1)
	for i := 0; i < sLen+1; i++ {
		dp[i] = make([]bool, pLen+1)
	}

	dp[0][0] = true
	for i := 1; i <= pLen; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if p[j-1] != '*' {
				// 单字符匹配并且前面的字符也需要匹配上
				dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '?')
			} else {
				// 当 p[j-1] == '*' 时
				//   要么，dp[i-1][j] == true，意味着，
				//   当 s[:i] 与 p[:j+1] 匹配，且p[j] == '*' 的时候，
				//   s[:i] 后接任意字符串的 s[:i+1] 仍与 p[:j+1] 匹配。
				//   要么，dp[i][j-1] == true，意味着，
				//   当 s[:i+1] 与 p[:j] 匹配后
				//   在 p[:j] 后添加'*'，s[:i+1] 与 p[:j+1] 任然匹配
				//   因为， '*' 可以匹配空字符。
				dp[i][j] = dp[i-1][j] || dp[i][j-1] // *匹配上空格；*匹配上一串字符
			}
		}
	}

	return dp[sLen][pLen]
}
