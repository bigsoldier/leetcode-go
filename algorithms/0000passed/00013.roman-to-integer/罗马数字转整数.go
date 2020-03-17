package code

func romanToInt(s string) int {
	var result int
	var preNum = getValue(s[0])
	for i := 1; i < len(s); i++ {
		num := getValue(s[i])
		if preNum < num {
			result -= preNum
		} else {
			result += preNum
		}
		preNum = num
	}
	result += preNum
	return result
}

func getValue(bt byte) int {
	switch bt {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
