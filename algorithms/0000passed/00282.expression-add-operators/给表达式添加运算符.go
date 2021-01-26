package code

func addOperators(num string, target int) (ans []string) {
	var dfs func(s, currStr string, curr, prefix int)
	dfs = func(s, currStr string, curr, prefix int) {
		if len(s) == 0 {
			if curr == target {
				ans = append(ans, currStr)
			}
			return
		}
		for i := 1; i <= len(s); i++ {
			left := s[:i]
			if left[0] == '0' && len(left) > 1 { // 以0开头的字符串无法变成数字
				return
			}
			leftNum := integer(left)
			right := s[i:]
			if len(currStr) == 0 {
				dfs(right, left, leftNum, leftNum)
			} else {
				dfs(right, currStr+"+"+left, curr+leftNum, leftNum)
				dfs(right, currStr+"-"+left, curr-leftNum, -leftNum)
				dfs(right, currStr+"*"+left, curr-prefix+prefix*leftNum, prefix*leftNum)
			}
		}
	}
	dfs(num, "", 0, 0)
	return
}

func integer(s string) int {
	res := int(s[0] - '0')
	for i := 1; i < len(s); i++ {
		res = res*10 + int(s[i]-'0')
	}
	return res
}
