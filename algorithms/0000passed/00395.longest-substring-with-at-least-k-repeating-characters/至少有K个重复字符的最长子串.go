package code

func longestSubstring(s string, k int) (ans int) {
	if len(s) == 0 {
		return 0
	}
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	var ch byte
	for i, c := range cnt {
		if c > 0 && c < k {
			ch = 'a' + byte(i)
			break
		}
	}
	if ch == 0 {
		return len(s)
	}
	for _, subStr := range strings.Split(s, string(ch)) {
		ans = max(ans, longestSubstring(subStr, k))
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
