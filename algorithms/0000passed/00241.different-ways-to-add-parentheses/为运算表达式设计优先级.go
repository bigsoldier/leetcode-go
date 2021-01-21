package code

import "strconv"

func diffWaysToCompute(input string) (ans []int) {
	if isDigit(input) {
		t, _ := strconv.Atoi(input)
		return []int{t}
	}
	for i, ch := range input {
		if ch == '+' || ch == '-' || ch == '*' {
			left := diffWaysToCompute(input[:i])
			right := diffWaysToCompute(input[i+1:])
			for _, l := range left {
				for _, r := range right {
					ans = append(ans, operate(l, r, byte(ch)))
				}
			}
		}
	}
	return
}

func isDigit(a string) bool {
	_, err := strconv.Atoi(a)
	if err != nil {
		return false
	}
	return true
}

func operate(a, b int, opt byte) int {
	switch opt {
	case '+':
		return a + b
	case '-':
		return a - b
	default:
		return a * b
	}
}
