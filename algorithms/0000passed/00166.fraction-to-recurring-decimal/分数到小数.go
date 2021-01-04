package code

func fractionToDecimal(numerator int, denominator int) string {
	sig := ""
	if numerator != 0 && (numerator&(1<<31)^(denominator&(1<<31)) > 0) {
		// int 默认是32位的,通过位运算计算符号大小
		// 如果通过乘积比较符号大小可能会导致越界
		sig = "-"
	}
	numerator, denominator = abs(numerator), abs(denominator)
	integer := strconv.Itoa(numerator / denominator) // 整数部分
	numerator = numerator % denominator * 10
	if numerator == 0 { // 没有小数部分
		return sig + integer
	}
	demical := "" // 小数部分
	m := map[string]int{}
	for i := 0; numerator > 0; i++ {
		key := strconv.Itoa(numerator) + "/" + strconv.Itoa(denominator)
		if v, ok := m[key]; ok { // 进入循环小数
			demical = demical[:v] + "(" + demical[v:i] + ")"
			break
		}
		m[key] = i
		demical += strconv.Itoa(numerator / denominator)
		numerator = numerator % denominator * 10
	}
	return sig + integer + "." + demical
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
