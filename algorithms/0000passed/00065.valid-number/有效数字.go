package code

// 判断是否是数字
func isNumber(s string) bool {
	s = trim(s) // 去除首尾的空格
	if len(s) == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		return isUnsignedNumber(s[1:], false)
	}
	return isUnsignedNumber(s, false)
}

// 是否是无符号实数，最多只能有一个逗号
func isUnsignedNumber(s string, hasDot bool) bool {
	if len(s) == 0 {
		return false
	}

	for i, c := range s {
		switch {
		case '0' <= c && c <= '9':
			continue
		case c == '.':
			if hasDot {
				return false
			}
			if i == len(s)-1 && i != 0 {
				// 以'.'结尾
				return true
			}
			if i+1 < len(s) && s[i+1] == 'e' {
				// 2.e3是正确的
				return i != 0 && isInterger(s[i+2:])
			}
			return isUnsignedNumber(s[i+1:], true)
		case c == 'e':
			if i == 0 {
				// e不能开头
				return false
			}
			return isInterger(s[i+1:])
		default:
			return false
		}
	}
	return true
}

func isInterger(s string) bool {
	if len(s) == 0 {
		return false
	}
	// 去符号
	if s[0] == '+' || s[0] == '-' {
		return isUnsignedInterger(s[1:])
	}
	return isUnsignedInterger(s)
}

func isUnsignedInterger(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		switch {
		case c < '0' || c > '9':
			return false
		}
	}
	return true
}

func trim(s string) string {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	j := len(s) - 1
	for i <= j && s[j] == ' ' {
		j--
	}

	return s[i : j+1]
}
