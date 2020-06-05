package code

func restoreIpAddresses(s string) []string {
	var result []string
	restoreIpAddress(nil, s, &result, 0)
	return result
}

func restoreIpAddress(ans []string, s string, result *[]string, depth int) {
	if depth == 4 {
		if len(s) != 0 {
			return
		}
		*result = append(*result, strings.Join(ans, "."))
		return
	}
	for i := 1; i < 4; i++ {
		if i <= len(s) && valid(s[:i]) {
			restoreIpAddress(append(ans, s[:i]), s[i:], result, depth+1)
		}
	}
}

// 判断字符串是否在0-255之间
func valid(s string) bool {
	switch len(s) {
	case 1, 2:
		if s[0] == '0' && len(s) != 1 {
			return false
		}
		return true
	case 3:
		if s[0] == '0' && len(s) != 1 {
			return false
		}
		if s < "256" {
			return true
		}
	}
	return false
}
