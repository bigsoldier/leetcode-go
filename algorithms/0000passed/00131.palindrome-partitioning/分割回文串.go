package code

func partition(s string) [][]string {
	var result [][]string
	backtracking(s, 0, len(s), []string{}, &result)
	return result
}

func backtracking(str string, start, end int, part []string, result *[][]string) {
	if start == end {
		s := make([]string, len(part))
		copy(s, part)
		*result = append(*result, s)
		return
	}
	for i := start; i < end; i++ {
		if !checkPalindrome(str, start, i) {
			continue
		}
		backtracking(str, i+1, end, append(part, str[start:i+1]), result)
	}
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
