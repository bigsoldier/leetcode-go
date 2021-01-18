package code

func calculate(s string) (sum int) {
	stack := []int{}
	num := 0
	var lastSign byte = '+'
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' || i == len(s)-1 {
			if lastSign == '+' {
				stack = append(stack, num)
			} else if lastSign == '-' {
				stack = append(stack, -num)
			} else if lastSign == '*' {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top*num)
			} else if lastSign == '/' {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top/num)
			}
			num = 0
			lastSign = s[i]
		}
	}
	for _, i := range stack {
		sum += i
	}
	return sum
}
