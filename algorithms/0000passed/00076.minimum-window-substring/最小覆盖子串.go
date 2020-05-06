package code

func minWindow(s string, t string) string {
	used := [128]int{}
	dict := [128]int{}
	for i := range t {
		dict[t[i]]++
	}

	size, total := len(s), len(t)

	min := size + 1
	res := ""

	// s[left:right+1] 就是 window
	// count 用于统计已有的 t 中字母的数量(在used中)。
	// count == total 表示已经收集完需要的全部字母
	for left, right, count := 0, 0, 0; right < size; right++ {
		if used[s[right]] < dict[s[right]] {
			// 出现了 window 中缺失的字母
			count++
		}
		used[s[right]]++

		// 保证 window 不丢失所需字母的前提下
		// 让 left 尽可能的大
		for left <= right && used[s[left]] > dict[s[left]] {
			used[s[left]]--
			left++
		}

		width := right - left + 1
		if count == total && min > width {
			min = width
			res = s[left : right+1]
		}

	}

	return res
}
