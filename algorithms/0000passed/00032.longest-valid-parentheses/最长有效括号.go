package code

func longestValidParentheses(s string) int {
	// 左右横跳
	// 从左向右遍历，从右向左遍历
	var left, right int
	var max int
	for _, v := range s {
		switch v {
		case '(':
			left++
		case ')':
			right++
		}

		if right > left {
			left, right = 0, 0
		} else if left == right {
			if left+right > max {
				max = left + right
			}
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '(':
			left++
		case ')':
			right++
		}
		if left > right {
			left, right = 0, 0
		} else if left == right {
			if left+right > max {
				max = left + right
			}
		}
	}

	return max
}
