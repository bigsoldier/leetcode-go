package code

func wordBreak(s string, wordDict []string) bool {
	var dict = make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}
	var dp = make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := dict[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
