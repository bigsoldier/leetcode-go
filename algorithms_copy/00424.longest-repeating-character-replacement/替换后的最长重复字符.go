package code

func characterReplacement(s string, k int) int {
	var left, maxCnt int
	var cnt [26]int
	for right, str := range s {
		cnt[str-'A']++
		maxCnt = max(maxCnt, cnt[str-'A'])
		if right-left+1-maxCnt > k {
			cnt[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
