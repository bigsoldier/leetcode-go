package code

func lengthOfLastWord(s string) int {
	// 去除右边的空格
	//s = strings.TrimRight(s," ") // 等价于下面的代码
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			s = s[:i+1]
			break
		}
	}
	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		count++
	}
	return count
}
